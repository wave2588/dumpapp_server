package controller

import "context"

type EmailController interface {
	SendRegisterEmail(ctx context.Context, title, content, receiveEmail string) error
	SendEmail(ctx context.Context, title, content, receiveEmail string, images []string) error
}
