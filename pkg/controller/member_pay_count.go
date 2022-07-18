package controller

import (
	"context"

	"dumpapp_server/pkg/common/enum"
)

type MemberPayCountController interface {
	AddCount(ctx context.Context, memberID, count int64, source enum.MemberPayCountSource) error
	CheckPayCount(ctx context.Context, loginID, limit int64) error
	DeductPayCount(ctx context.Context, loginID, limit int64, status enum.MemberPayCountStatus, use enum.MemberPayCountUse) error
}
