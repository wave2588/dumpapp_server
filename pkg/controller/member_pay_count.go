package controller

import (
	"context"

	"dumpapp_server/pkg/common/enum"
)

type MemberDownloadController interface {
	CheckPayCount(ctx context.Context, loginID, limit int64) error
	DeductPayCount(ctx context.Context, loginID, limit int64, use enum.MemberPayCountUse) error
}
