package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/datatype"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/util"
	pkgErr "github.com/pkg/errors"
	"github.com/spf13/cast"
)

type AppSourceController struct {
	appSourceDAO       dao.AppSourceDAO
	memberAppSourceDAO dao.MemberAppSourceDAO
}

var DefaultAppSourceController *AppSourceController

func init() {
	DefaultAppSourceController = NewAppSourceController()
}

func NewAppSourceController() *AppSourceController {
	return &AppSourceController{
		appSourceDAO:       impl.DefaultAppSourceDAO,
		memberAppSourceDAO: impl.DefaultMemberAppSourceDAO,
	}
}

func (c *AppSourceController) Insert(ctx context.Context, loginID int64, URL string) (int64, error) {
	/// 检查 URL 是否有效
	appSourceInfoMap, err := c.BatchGetAppSourceInfo(ctx, []string{URL})
	if err != nil {
		return 0, err
	}
	_, ok := appSourceInfoMap[URL]
	if !ok {
		return 0, errors.ErrAppSourceDisabled
	}

	var appSourceID int64
	appSourceMap, err := c.appSourceDAO.BatchGetByURL(ctx, []string{URL})
	if err != nil {
		return 0, err
	}
	appSource := appSourceMap[URL]
	if appSource == nil {
		appSourceID = util.MustGenerateID(ctx)
		err = c.appSourceDAO.Insert(ctx, &models.AppSource{
			ID:     appSourceID,
			URL:    URL,
			BizExt: datatype.AppSourceBizExt{},
		})
	} else {
		appSourceID = appSource.ID
	}

	memberAppSource, err := c.memberAppSourceDAO.GetByMemberIDAppSourceID(ctx, loginID, appSourceID)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		return 0, err
	}
	if memberAppSource != nil {
		return memberAppSource.ID, nil
	}

	/// 如果没绑定过则写入
	memberAppSourceID := util.MustGenerateID(ctx)
	err = c.memberAppSourceDAO.Insert(ctx, &models.MemberAppSource{
		ID:          memberAppSourceID,
		MemberID:    loginID,
		AppSourceID: appSourceID,
	})
	return memberAppSourceID, nil
}

func (c *AppSourceController) BatchGetAppSourceInfo(ctx context.Context, URLs []string) (map[string]*controller.AppSourceInfo, error) {
	udid := cast.ToString(ctx.Value(constant.CtxKeyAppUDID))
	result := make(map[string]*controller.AppSourceInfo)
	batch := util.NewBatch(ctx)
	for idx, URL := range URLs {
		batch.Append(func(idx int, U string) util.FutureFunc {
			return func() error {
				URI := fmt.Sprintf("%s?udid=%s", U, udid)
				resp, err := util.HttpRequest("GET", URI, map[string]string{}, map[string]string{}, 10*time.Second)
				if err != nil {
					return err
				}
				var info *controller.AppSourceInfo
				err = json.Unmarshal(resp, &info)
				if err != nil {
					return err
				}
				result[U] = info
				return nil
			}
		}(idx, URL))
	}
	batch.Wait()
	return result, nil
}
