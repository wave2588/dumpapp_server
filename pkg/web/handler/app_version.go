package handler

import (
	"context"
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"github.com/Masterminds/semver/v3"
)

type AppVersionHandler struct {
	searchRecordV2DAO dao.SearchRecordV2DAO
	dumpappVersionDAO dao.DumpappVersionDAO
}

func NewAppVersionHandler() *AppVersionHandler {
	return &AppVersionHandler{
		searchRecordV2DAO: impl.DefaultSearchRecordV2DAO,
		dumpappVersionDAO: impl.DefaultDumpappVersionDAO,
	}
}

type result struct {
	IsNeedUpdate  bool   `json:"is_need_update"`  /// 是否需要更新
	IsForceUpdate bool   `json:"is_force_update"` /// 是否需要强制更新
	DownloadURL   string `json:"download_url"`    /// 下载地址
}

func (h *AppVersionHandler) CheckAppVersion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	appVersion, ok := ctx.Value(constant.CtxKeyAppVersion).(string)
	if !ok {
		panic(errors.HttpUnprocessableError)
	}
	if appVersion == "" {
		panic(errors.UnproccessableError("缺少版本号字段"))
	}

	lastDumpappVersion, err := h.getLastDumpappVersion(ctx)
	util.PanicIf(err)

	constrain, err := semver.NewConstraint(fmt.Sprintf("<%s", lastDumpappVersion.Version)) /// 需要更新的版本
	util.PanicIf(err)

	v, err := semver.NewVersion(appVersion)
	util.PanicIf(err)

	result := &result{
		IsNeedUpdate:  false,
		IsForceUpdate: lastDumpappVersion.IsForceUpdate,
		DownloadURL:   "https://www.baidu.com",
	}

	if constrain.Check(v) {
		result.IsNeedUpdate = true
	}
	util.RenderJSON(w, result)
}

func (h *AppVersionHandler) getLastDumpappVersion(ctx context.Context) (*models.DumpappVersion, error) {
	ids, err := h.dumpappVersionDAO.ListIDs(ctx, 0, 1, nil, []string{"id desc"})
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors.UnproccessableError("未找到最新版本")
	}
	resMap, err := h.dumpappVersionDAO.BatchGet(ctx, ids)
	if err != nil {
		return nil, err
	}
	return resMap[ids[0]], nil
}
