package controller

import (
	"context"
)

type EmailWebController interface {
	SendEmailToMaster(ctx context.Context, appName, receiveEmail string) error
}
