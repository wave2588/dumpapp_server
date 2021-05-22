package handler

import (
	"context"
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
}

func NewAdminIpaHandler() *AdminIpaHandler {
	return &AdminIpaHandler{
		ipaDAO:        impl.DefaultIpaDAO,
		ipaVersionDAO: impl.DefaultIpaVersionDAO,
	}
}

type createIpaArgs struct {
	Ipas []*ipaArgs `json:"ipas" validate:"required"`
}

type ipaArgs struct {
	Name    string `json:"name" validate:"required"`
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

	names := make([]string, 0)
	for _, ipa := range args.Ipas {
		names = append(names, ipa.Name)
	}

	ipaMap, err := h.ipaDAO.BatchGetByName(ctx, names)
	util.PanicIf(err)

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	for _, ipaArgs := range args.Ipas {
		ipa := ipaMap[ipaArgs.Name]
		var ipaID int64
		if ipa == nil { /// 没找到说明不存在, 往库里写入
			util.PanicIf(h.ipaDAO.Insert(ctx, &models.Ipa{
				Name: ipaArgs.Name,
			}))
			/// 再获取一遍是为了得到 id
			i, err := h.ipaDAO.GetByName(ctx, ipaArgs.Name)
			util.PanicIf(err)
			ipaID = i.ID
		} else {
			ipaID = ipa.ID
		}
		/// fixme: 需要判断 版本 是否存在, 存在的话就不能再保存.
		util.PanicIf(h.ipaVersionDAO.Insert(ctx, &models.IpaVersion{
			IpaID:     ipaID,
			Version:   ipaArgs.Version,
			TokenPath: ipaArgs.Token,
		}))

		/// todo: 后期如果做 ipa 个数限制的话, 在这里做.
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	util.RenderJSON(w, "保存成功")
}
