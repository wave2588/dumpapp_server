package controller

import (
	"context"
)

type EmailWebController interface {
	SendEmailToMaster(ctx context.Context, appName, version, memberEmail string) error
	SendVipEmailToMaster(ctx context.Context, appName, version, memberEmail string) error
	SendUpdateIpaEmail(ctx context.Context, ipaID int64, email, name string) error
}
