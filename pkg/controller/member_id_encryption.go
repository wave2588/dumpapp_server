package controller

import "context"

type MemberIDEncryptionController interface {
	GetCodeByMemberID(ctx context.Context, memberID int64) (string, error)
	GetMemberIDByCode(ctx context.Context, code string) (int64, error)
}
