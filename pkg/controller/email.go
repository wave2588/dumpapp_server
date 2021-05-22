package controller

import "context"

type EmailController interface {
	SendEmail(ctx context.Context, title, content, receiveEmail string, images []string) error
}
