package controller

import (
	"context"
)

type IpaSignWebController interface {
	Sign(ctx context.Context, loginID, certificateID, ipaVersionID int64) error
}
