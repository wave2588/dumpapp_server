package controller

import "context"

type CertificateV2WebController interface {
	Create(ctx context.Context, loginID int64, UDID, note string, priceID int64) (int64, error)

	// 仅仅给管理后台使用
	AdminCreate(ctx context.Context, memberID int64, UDID string) (int64, error)
}
