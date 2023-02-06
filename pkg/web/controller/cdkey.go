package controller

import "context"

type CdkeyWebController interface {
	Create(ctx context.Context, memberID int64) (int64, error)
}
