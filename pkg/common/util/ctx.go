package util

import (
	"context"
	"errors"
	"fmt"
)

func GetStringCtxValue(ctx context.Context, key string) string {
	if v := ctx.Value(key); v != nil {
		switch v.(type) {
		case string:
			return v.(string)
		case int:
			return fmt.Sprintf("%d", v.(int))
		case int64:
			return fmt.Sprintf("%d", v.(int64))
		default:
			panic(errors.New("not support type"))
		}
	}
	return ""
}

type contextKey struct {
	Name string
}

var requestCtxKey = &contextKey{"context"}

func GetRequestContext(ctx context.Context, key string) string {
	if v := ctx.Value(requestCtxKey); v != nil {
		return v.(map[string]string)[key]
	} else {
		return ""
	}
}

func ResetCtxKey(ctx context.Context, key interface{}) context.Context {
	return context.WithValue(ctx, key, nil)
}
