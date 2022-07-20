package controller

import "context"

type CertificateWebController interface {
	PayCertificate(ctx context.Context, loginID int64, udid string, payCount int64, payType string) (int64, error)
	GetModifiedCertificateData(ctx context.Context, p12Data, originalPassword, newPassword string) (string, error)
}
