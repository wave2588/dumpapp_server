package controller

import (
	"context"

	"dumpapp_server/pkg/common/enum"
)

type DispenseCountController interface {
	AddCount(ctx context.Context, memberID, count int64) error
	Check(ctx context.Context, memberID, count int64) error
	DeductCount(ctx context.Context, memberID, count int64, recordType enum.DispenseCountRecordType) error
}
