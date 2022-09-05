package controller

import (
	"context"

	"dumpapp_server/pkg/common/enum"
)

type DispenseCountController interface {
	AddCount(ctx context.Context, memberID, count int64, recordType enum.DispenseCountRecordType) error
	Check(ctx context.Context, memberID, count int64) error
	DeductCount(ctx context.Context, memberID, count int64, recordType enum.DispenseCountRecordType) error

	/// 计算 member_sign_ipa 需要消耗的分发卷次数
	CalculateMemberSignIpaDispenseCount(ctx context.Context, ipaSize int64) int64
}
