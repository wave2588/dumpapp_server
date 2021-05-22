package handler

import (
	"context"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"github.com/go-playground/validator/v10"
)

type AdminIpaHandler struct {
	ipaDAO        dao.IpaDAO
	ipaVersionDAO dao.IpaVersionDAO

	appleCtl controller.AppleController
}

func NewAdminIpaHandler() *AdminIpaHandler {
	return &AdminIpaHandler{
		ipaDAO:        impl.DefaultIpaDAO,
		ipaVersionDAO: impl.DefaultIpaVersionDAO,

		appleCtl: impl2.DefaultAppleController,
	}
}

type createIpaArgs struct {
	Ipas []*ipaArgs `json:"ipas" validate:"required"`
}

type ipaArgs struct {
	AppID   int64  `json:"app_id" validate:"required"`
	Version string `json:"version" validate:"required"`
	Token   string `json:"token" validate:"required"`
}

func (p *createIpaArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminIpaHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &createIpaArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	appIDs := make([]int64, 0)
	appInfoMap := make([]*controller.AppInfo, 0)
	for _, ipa := range args.Ipas {
		appInfo, err := h.appleCtl.GetAppInfoByAppID(ctx, ipa.AppID)
		util.PanicIf(err)
		appInfoMap[ipa.AppID] = appInfo
		appIDs = append(appIDs, ipa.AppID)
	}

	ipaMap, err := h.ipaDAO.BatchGet(ctx, appIDs)
	util.PanicIf(err)

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	for _, ipaArgs := range args.Ipas {
		appInfo := appInfoMap[ipaArgs.AppID]
		if appInfo == nil {
			continue
		}
		ipa := ipaMap[ipaArgs.AppID]
		if ipa == nil {
			util.PanicIf(h.ipaDAO.Insert(ctx, &models.Ipa{
				Name: appInfo.Name,
			}))
		}
		/// todo: 后期如果做 ipa 个数限制的话, 在这里做.
		util.PanicIf(h.ipaVersionDAO.Insert(ctx, &models.IpaVersion{
			IpaID:     ipaArgs.AppID,
			Version:   ipaArgs.Version,
			TokenPath: ipaArgs.Token,
		}))
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	util.RenderJSON(w, "保存成功")
}
