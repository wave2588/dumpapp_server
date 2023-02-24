package controller

import "context"

type CdKeyController interface {
	AddCdKeyByMemberBuyCertificate(ctx context.Context, certificateID int64) error
}
