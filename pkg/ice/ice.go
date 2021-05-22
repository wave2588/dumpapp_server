package rpc

import (
	"context"
)

type IceRPC interface {
	MustGenerateID(ctx context.Context) int64
	MustGenerateCaptcha(ctx context.Context) string
}
