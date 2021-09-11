package controller

import (
	"context"
)

type EmailWebController interface {
	SendUpdateIpaEmail(ctx context.Context, ipaID int64, email, name string) error
}
