package controller

import "context"

type AuthController interface {
	Auth(ctx context.Context)
}
