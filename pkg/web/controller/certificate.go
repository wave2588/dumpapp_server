package controller

import "context"

type CertificateWebController interface {
	GetModifiedCertificateData(ctx context.Context, p12Data string) (string, error)
}
