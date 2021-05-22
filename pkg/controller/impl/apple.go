package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"dumpapp_server/pkg/controller"
	errors2 "dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/util"
	"github.com/spf13/cast"
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
	ResultCount int64     `json:"resultCount"`
	Results     []*result `json:"results"`
}

type result struct {
	TrackId        int64  `json:"trackId"`
	TrackName      string `json:"trackName"`
	FormattedPrice string `json:"formattedPrice"`
	BundleID       string `json:"bundleId"`
}

func (c *AppleController) GetAppInfoByAppID(ctx context.Context, appID int64) (*controller.AppInfo, error) {
	res, err := http.Get(fmt.Sprintf("http://itunes.apple.com/cn/lookup?id=%d", appID))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	app := &appResult{}
	err = json.Unmarshal(body, app)
	if err != nil {
		return nil, err
	}

	if len(app.Results) == 0 {
		return nil, errors2.ErrNotFoundApp
	}

	r := app.Results[0]
	price := strings.ReplaceAll(r.FormattedPrice, "¥", "")
	return &controller.AppInfo{
		AppID:    r.TrackId,
		Name:     r.TrackName,
		BundleID: r.BundleID,
		Price:    cast.ToFloat64(price),
	}, nil
}

func (c *AppleController) BatchGetAppInfoByBundleIDs(ctx context.Context, bundleInfos []*controller.BundleInfo) (map[string]*controller.AppInfo, error) {
	res := make([]*controller.AppInfo, 0)
	batch := util.NewBatch(ctx)
	for _, bundleInfo := range bundleInfos {
		batch.Append(func(bundleInfo *controller.BundleInfo) util.FutureFunc {
			return func() error {
				appInfo, err := c.GetAppInfoByBundleID(ctx, bundleInfo.BundleID, bundleInfo.IsDomestic)
				if err != nil {
					return err
				}
				res = append(res, appInfo)
				return nil
			}
		}(bundleInfo))
	}
	rpcErrs := batch.Get()
	result := make(map[string]*controller.AppInfo)
	for idx, bundleInfo := range bundleInfos {
		if rpcErrs[idx] != nil {
			err := rpcErrs[idx]
			return nil, err
		}
		result[bundleInfo.BundleID] = res[idx]
	}
	return result, nil
}

func (c *AppleController) GetAppInfoByBundleID(ctx context.Context, bundleID string, isDomestic bool) (*controller.AppInfo, error) {
	url := fmt.Sprintf("http://itunes.apple.com/lookup?bundleId=%s", bundleID)
	if isDomestic {
		url = fmt.Sprintf("http://itunes.apple.com/cn/lookup?bundleId=%s", bundleID)
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	app := &appResult{}
	err = json.Unmarshal(body, app)
	if err != nil {
		return nil, err
	}
	if len(app.Results) == 0 {
		return nil, errors2.ErrNotFoundApp
	}

	r := app.Results[0]
	price := strings.ReplaceAll(r.FormattedPrice, "¥", "")
	return &controller.AppInfo{
		AppID:    r.TrackId,
		Name:     r.TrackName,
		BundleID: r.BundleID,
		Price:    cast.ToFloat64(price),
	}, nil
}
