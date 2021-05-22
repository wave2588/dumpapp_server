package controller

import "context"

type AppleController interface {
	GetAppInfoByAppID(ctx context.Context, appID int64) (*AppInfo, error)
}

type AppInfo struct {
	AppID int64   `json:"app_id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
