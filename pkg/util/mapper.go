package util

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	pkgErrors "github.com/pkg/errors"
)

type MapperResult struct {
	Input  interface{}
	Output interface{}
	Err    error
}

func ConcurrentMapper(ctx context.Context, inputs []interface{}, mapper func(input interface{}) (interface{}, error), mapperTimeout time.Duration) (response chan *MapperResult) {
	response = make(chan *MapperResult)
	go func() {
		defer close(response)
		ctxTmp, cancel := context.WithTimeout(ctx, mapperTimeout)
		defer cancel()
		// 业务逻辑处理
		batch := NewBatch(ctxTmp)
		for idx, input := range inputs {
			batch.Append(func(idx int64, input interface{}) FutureFunc {
				return func() error {
					res, err := mapper(input)
					response <- &MapperResult{
						Input:  input,
						Output: res,
					}
					return err
				}
			}(int64(idx), input))
		}
		// 返回后续的 Err
		resErr := batch.Get()
		for _, err := range resErr {
			if err != nil {
				response <- &MapperResult{
					Err: err,
				}
			}
		}
	}()
	return response
}

type Future struct {
	err error
	ch  chan struct{}
}

type FutureFunc func() error

type Batch struct {
	ctx     context.Context
	futures []*Future
	wg      *sync.WaitGroup
}

func NewBatch(ctx context.Context) *Batch {
	return &Batch{
		ctx: ctx,
		wg:  &sync.WaitGroup{},
	}
}

func (b *Batch) Append(fn FutureFunc) {
	b.wg.Add(1)

	f := &Future{
		ch: make(chan struct{}),
	}
	b.futures = append(b.futures, f)
	go func() {
		defer func() {
			if rval := recover(); rval != nil {
				if err, ok := rval.(error); ok {
					f.err = pkgErrors.WithStack(err)
				} else {
					rvalStr := fmt.Sprint(rval)
					f.err = pkgErrors.WithStack(errors.New(rvalStr))
				}
			}
			close(f.ch)
			b.wg.Done()
		}()

		// 执行结果
		f.err = fn()
	}()
}

func (b *Batch) Get() []error {
	b.Wait()
	errs := make([]error, len(b.futures))
	for idx, f := range b.futures {
		var tmpErr error
		select {
		case <-f.ch:
			tmpErr = f.err
		default:
		}
		if tmpErr != nil {
			errs[idx] = tmpErr
		} else {
			errs[idx] = b.ctx.Err()
		}

	}
	return errs
}

func (b *Batch) Wait() {
	select {
	case <-b.ctx.Done():
		return
	default:
	}
	b.wg.Wait()
}
