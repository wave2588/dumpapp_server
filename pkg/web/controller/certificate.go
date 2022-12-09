package controller

import "context"

type CertificateWebController interface {
	PayCertificate(ctx context.Context, loginID int64, udid, note string, priceID int64, isReplenish bool, payType string) (int64, error)
	GetModifiedCertificateData(ctx context.Context, p12Data, originalPassword, newPassword string) (string, error)
	CertificateReplenish(ctx context.Context, loginID, cerID int64) (int64, error)
}
