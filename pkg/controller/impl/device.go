package impl

import (
	"context"
)

type DeviceController struct{}

var DefaultDeviceController *DeviceController

func init() {
	DefaultDeviceController = NewDeviceController()
}

func NewDeviceController() *DeviceController {
	return &DeviceController{}
}

func (c *DeviceController) GetConfigQRCode(ctx context.Context, memberID int64) {
}
