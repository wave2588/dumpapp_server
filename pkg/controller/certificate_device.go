package controller

import "context"

type CertificateDeviceController interface {
	IsReplenish(ctx context.Context, memberID int64, UDID string) (bool, error)
}
