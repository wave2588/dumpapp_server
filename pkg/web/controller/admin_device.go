package controller

import "context"

type AdminDeviceController interface {
	Unbind(ctx context.Context, email, udid string) error
}
