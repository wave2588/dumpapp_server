package controller

import "context"

type CertificateCreateWebController interface {
	Create(ctx context.Context, loginID int64, udid, note string, priceID int64, isReplenish bool, payType string) (int64, error)
}
