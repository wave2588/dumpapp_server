package controller

import (
	"context"
)

type EmailWebController interface {
	SendEmailToMaster(ctx context.Context, appName, version, memberEmail string) error
	SendUpdateIpaEmail(ctx context.Context, email, name string) error
}
