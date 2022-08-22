package controller

import (
	"context"

	"dumpapp_server/pkg/common/enum"
)

type MemberSignIpaDownloadCountController interface {
	CheckCount(ctx context.Context, memberID, limit int64) (bool, error)
	AddCount(ctx context.Context, memberID, count int64, recordType enum.MemberSignIpaDownloadCountRecordType) error
	DeductPayCount(ctx context.Context, memberID, count int64, recordType enum.MemberSignIpaDownloadCountRecordType) error
}
