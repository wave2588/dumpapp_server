package controller

import "context"

type AppleController interface {
	GetAppInfoByAppID(ctx context.Context, appID int64) (*AppInfo, error)
	BatchGetAppInfoByAppIDs(ctx context.Context, appIDs []int64) (map[int64]*AppInfo, error)
	GetAppInfoByBundleID(ctx context.Context, bundleID string, isDomestic bool) (*AppInfo, error)
	BatchGetAppInfoByBundleIDs(ctx context.Context, bundleInfos []*BundleInfo) (map[string]*AppInfo, error)
}

type AppInfo struct {
	AppID    int64   `json:"app_id"`
	Name     string  `json:"name"`
	BundleID string  `json:"bundle_id"`
	Price    float64 `json:"price"`
}

type BundleInfo struct {
	BundleID   string
	IsDomestic bool
}
