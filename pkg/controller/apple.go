package controller

import "context"

type AppleController interface {
	GetAppInfoByAppID(ctx context.Context, appID int64) (*AppInfo, error)
	GetAppInfoByBundleID(ctx context.Context, bundleID string, isDomestic bool) (*AppInfo, error)
}

type AppInfo struct {
	AppID    int64   `json:"app_id"`
	Name     string  `json:"name"`
	BundleID string  `json:"bundle_id"`
	Price    float64 `json:"price"`
}
