package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"dumpapp_server/pkg/util"
)

type AppleController struct{}

var DefaultAppleController *AppleController

func init() {
	DefaultAppleController = NewAppleController()
}

func NewAppleController() *AppleController {
	return &AppleController{}
}

type appResult struct {
	ResultCount int64         `json:"resultCount"`
	Results     []interface{} `json:"results"`
}

func (c *AppleController) BatchGetAppInfoByAppIDs(ctx context.Context, appIDs []int64) (map[int64]interface{}, error) {
	res := make([]interface{}, len(appIDs))
	batch := util.NewBatch(ctx)
	for idx, appID := range appIDs {
		batch.Append(func(idx int, appID int64) util.FutureFunc {
			return func() error {
				appInfo, err := c.GetAppInfoByAppID(ctx, appID)
				if err != nil {
					return err
				}
				if appInfo == nil {
					return nil
				}
				res[idx] = appInfo
				return nil
			}
		}(idx, appID))
	}
	rpcErrs := batch.Get()
	result := make(map[int64]interface{})
	for idx, appID := range appIDs {
		if rpcErrs[idx] != nil {
			err := rpcErrs[idx]
			return nil, err
		}
		data := res[idx]
		if data == nil {
			continue
		}
		result[appID] = res[idx]
	}
	return result, nil
}

func (c *AppleController) GetAppInfoByAppID(ctx context.Context, appID int64) (interface{}, error) {
	url := fmt.Sprintf("https://itunes.apple.com/cn/lookup?id=%d", appID)
	resp, err := http.DefaultClient.Post(url, "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	app := &appResult{}
	err = json.Unmarshal(body, app)
	if err != nil {
		return nil, err
	}

	if len(app.Results) == 0 {
		return nil, nil
	}
	return app.Results[0], nil
}

//func (c *AppleController) BatchGetAppInfoByBundleIDs(ctx context.Context, bundleInfos []*controller.BundleInfo) (map[string]*controller.AppInfo, error) {
//	res := make(map[int]*controller.AppInfo)
//	batch := util.NewBatch(ctx)
//	for idx, bundleInfo := range bundleInfos {
//		batch.Append(func(idx int, bundleInfo *controller.BundleInfo) util.FutureFunc {
//			return func() error {
//				appInfo, err := c.GetAppInfoByBundleID(ctx, bundleInfo.BundleID, bundleInfo.IsDomestic)
//				if err != nil {
//					return err
//				}
//				res[idx] = appInfo
//				return nil
//			}
//		}(idx, bundleInfo))
//	}
//	rpcErrs := batch.Get()
//	result := make(map[string]*controller.AppInfo)
//	for idx := range bundleInfos {
//		if rpcErrs[idx] != nil {
//			err := rpcErrs[idx]
//			return nil, err
//		}
//		result[res[idx].BundleID] = res[idx]
//	}
//	return result, nil
//}

//func (c *AppleController) GetAppInfoByBundleID(ctx context.Context, bundleID string, isDomestic bool) (*controller.AppInfo, error) {
//	url := fmt.Sprintf("http://itunes.apple.com/lookup?bundleId=%s", bundleID)
//	if isDomestic {
//		url = fmt.Sprintf("http://itunes.apple.com/cn/lookup?bundleId=%s", bundleID)
//	}
//
//	res, err := http.Get(url)
//	if err != nil {
//		return nil, err
//	}
//	defer res.Body.Close()
//	body, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		return nil, err
//	}
//	app := &appResult{}
//	err = json.Unmarshal(body, app)
//	if err != nil {
//		return nil, err
//	}
//	if len(app.Results) == 0 {
//		return nil, errors2.ErrNotFoundApp
//	}
//
//	r := app.Results[0]
//	price := strings.ReplaceAll(r.FormattedPrice, "¥", "")
//	return &controller.AppInfo{
//		AppID:    r.TrackId,
//		Name:     r.TrackName,
//		BundleID: r.BundleID,
//		Price:    cast.ToFloat64(price),
//	}, nil
//}
