package util

import (
	"github.com/pkg/errors"
	"io"
)

func PanicIf(err error) {
	if err != nil {
		if err == io.EOF {
			return
		}
		e := errors.WithStack(err)
		panic(e)
	}
}
