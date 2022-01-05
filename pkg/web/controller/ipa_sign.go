package controller

import (
	"context"
)

type IpaSignWebController interface {
	AddSignTask(ctx context.Context, loginID, certificateID, ipaVersionID int64) error
}
