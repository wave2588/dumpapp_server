package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	controller2 "dumpapp_server/pkg/web/controller"
	impl3 "dumpapp_server/pkg/web/controller/impl"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AdminIpaHandler struct {
	accountDAO        dao.AccountDAO
	ipaDAO            dao.IpaDAO
	ipaVersionDAO     dao.IpaVersionDAO
	searchRecordV2DAO dao.SearchRecordV2DAO

	appleCtl    controller.AppleController
	emailWebCtl controller2.EmailWebController
}

func NewAdminIpaHandler() *AdminIpaHandler {
	return &AdminIpaHandler{
		accountDAO:        impl.DefaultAccountDAO,
		ipaDAO:            impl.DefaultIpaDAO,
		ipaVersionDAO:     impl.DefaultIpaVersionDAO,
		searchRecordV2DAO: impl.DefaultSearchRecordV2DAO,

		appleCtl:    impl2.DefaultAppleController,
		emailWebCtl: impl3.DefaultEmailWebController,
	}
}

type createIpaArgs struct {
	Ipas []*ipaArgs `json:"ipas" validate:"required"`
}

type ipaArgs struct {
	IpaID    string `json:"ipa_id" validate:"required"`
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

	ipaArgsMap := make(map[int64]*ipaArgs)
	ipaIDs := make([]int64, 0)
	for _, ipaArgs := range args.Ipas {
		ipaID := cast.ToInt64(ipaArgs.IpaID)
		ipaIDs = append(ipaIDs, ipaID)
		ipaArgsMap[ipaID] = ipaArgs
	}
	ipaMap, err := h.ipaDAO.BatchGet(ctx, ipaIDs)
	util.PanicIf(err)

	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	for _, ipaArgs := range args.Ipas {
		ipaID := cast.ToInt64(ipaArgs.IpaID)
		ipa := ipaMap[ipaID]
		if ipa == nil {
			util.PanicIf(h.ipaDAO.Insert(ctx, &models.Ipa{
				ID:       ipaID,
				Name:     ipaArgs.Name,
				BundleID: ipaArgs.BundleID,
			}))
		}
		/// todo: 后期如果做 ipa 个数限制的话, 在这里做.
		util.PanicIf(h.ipaVersionDAO.Insert(ctx, &models.IpaVersion{
			IpaID:     ipaID,
			Version:   ipaArgs.Version,
			TokenPath: ipaArgs.Token,
		}))
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	util.PanicIf(h.sendEmail(ctx, ipaArgsMap))

	util.RenderJSON(w, "保存成功")
}

func (h *AdminIpaHandler) sendEmail(ctx context.Context, ipaArgsMap map[int64]*ipaArgs) error {
	ipaIDs := make([]int64, 0)
	for _, args := range ipaArgsMap {
		ipaID := cast.ToInt64(args.IpaID)
		ipaIDs = append(ipaIDs, ipaID)
	}

	filters := []qm.QueryMod{
		models.SearchRecordV2Where.CreatedAt.GTE(time.Date(0, 0, -7, 0, 0, 0, 0, time.Now().Location())),
	}
	records, err := h.searchRecordV2DAO.BatchGetByIpaIDs(ctx, ipaIDs, filters)
	if err != nil {
		return err
	}
	memberIDs := make([]int64, 0)
	for _, record := range records {
		memberIDs = append(memberIDs, record.MemberID)
	}
	memberIDs = util2.RemoveDuplicates(memberIDs)
	memberMap, err := h.accountDAO.BatchGet(ctx, memberIDs)
	if err != nil {
		return err
	}

	// 已发过滤
	filterMap := make(map[string]struct{})

	batch := util2.NewBatch(ctx)
	for _, record := range records {
		member := memberMap[record.MemberID]
		if member == nil {
			continue
		}
		ipaArgs := ipaArgsMap[record.IpaID]
		if ipaArgs == nil {
			continue
		}
		key := fmt.Sprintf("%d-%s", member.ID, ipaArgs.IpaID)
		if _, ok := filterMap[key]; ok {
			continue
		}
		filterMap[key] = struct{}{}
		batch.Append(func() error {
			return h.emailWebCtl.SendUpdateIpaEmail(ctx, cast.ToInt64(ipaArgs.IpaID), member.Email, ipaArgs.Name)
		})
	}
	batch.Wait()

	return nil
}
