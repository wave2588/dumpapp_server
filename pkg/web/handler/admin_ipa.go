package handler

import (
	"context"
	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
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
	IpaID    int64  `json:"ipa_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	BundleID string `json:"bundle_id" validate:"required"`
	Version  string `json:"version" validate:"required"`
	Token    string `json:"token" validate:"required"`
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

	ipaIDs := make([]int64, 0)
	for _, ipaArgs := range args.Ipas {
		ipaIDs = append(ipaIDs, ipaArgs.IpaID)
	}
	ipaMap, err := h.ipaDAO.BatchGet(ctx, ipaIDs)
	util.PanicIf(err)

	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	for _, ipaArgs := range args.Ipas {
		ipa := ipaMap[ipaArgs.IpaID]
		if ipa == nil {
			util.PanicIf(h.ipaDAO.Insert(ctx, &models.Ipa{
				ID:       ipaArgs.IpaID,
				Name:     ipaArgs.Name,
				BundleID: ipaArgs.BundleID,
			}))
		}
		/// todo: 后期如果做 ipa 个数限制的话, 在这里做.
		util.PanicIf(h.ipaVersionDAO.Insert(ctx, &models.IpaVersion{
			IpaID:     ipaArgs.IpaID,
			Version:   ipaArgs.Version,
			TokenPath: ipaArgs.Token,
		}))
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	util.RenderJSON(w, "保存成功")
}
