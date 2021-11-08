package controller

import "context"

type CertificateWebController interface {
	ModifyCertificateContent(ctx context.Context, p12Data string) (string, error)
}
