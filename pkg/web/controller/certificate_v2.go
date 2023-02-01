package controller

import "context"

type CertificateV2WebController interface {
	Create(ctx context.Context, loginID int64, UDID, note string, priceID int64) (int64, error)
}
