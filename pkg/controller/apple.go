package controller

import "context"

type AppleController interface {
	GetAppInfoByAppID(ctx context.Context, appID int64) (interface{}, error)
	BatchGetAppInfoByAppIDs(ctx context.Context, appIDs []int64) (map[int64]interface{}, error)
	// GetAppInfoByBundleID(ctx context.Context, bundleID string, isDomestic bool) (*AppInfo, error)
	// BatchGetAppInfoByBundleIDs(ctx context.Context, bundleInfos []*BundleInfo) (map[string]*AppInfo, error)
}
