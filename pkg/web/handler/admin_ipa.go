package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	util2 "dumpapp_server/pkg/util"
	controller2 "dumpapp_server/pkg/web/controller"
	impl3 "dumpapp_server/pkg/web/controller/impl"
	"github.com/go-playground/validator/v10"
	pkgErr "github.com/pkg/errors"
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
	tencentCtl  controller.TencentController
}

func NewAdminIpaHandler() *AdminIpaHandler {
	return &AdminIpaHandler{
		accountDAO:        impl.DefaultAccountDAO,
		ipaDAO:            impl.DefaultIpaDAO,
		ipaVersionDAO:     impl.DefaultIpaVersionDAO,
		searchRecordV2DAO: impl.DefaultSearchRecordV2DAO,

		appleCtl:    impl2.DefaultAppleController,
		emailWebCtl: impl3.DefaultEmailWebController,
		tencentCtl:  impl2.DefaultTencentController,
	}
}

type createIpaArgs struct {
	Ipas []*ipaArgs `json:"ipas" validate:"required"`
}

type ipaArgs struct {
	IpaID     string     `json:"ipa_id" validate:"required"`
	Name      string     `json:"name" validate:"required"`
	BundleID  string     `json:"bundle_id" validate:"required"`
	IsInterim bool       `json:"is_interim"`
	Versions  []*Version `json:"versions"`
}

type Version struct {
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
	loginID := middleware.MustGetMemberID(ctx)
	if _, ok := constant.OpsAuthMemberIDMap[loginID]; !ok {
		panic(errors.ErrMemberAccessDenied)
	}

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

	ipaVersionMap, err := h.ipaVersionDAO.BatchGetIpaVersions(ctx, ipaIDs)
	util.PanicIf(err)

	for _, ipaArgs := range args.Ipas {
		ipaID := cast.ToInt64(ipaArgs.IpaID)
		ipa := ipaMap[ipaID]
		if ipa == nil {
			util.PanicIf(h.ipaDAO.Insert(ctx, &models.Ipa{
				ID:        ipaID,
				Name:      ipaArgs.Name,
				BundleID:  ipaArgs.BundleID,
				IsInterim: cast.ToInt(ipaArgs.IsInterim),
			}))
		} else {
			ipa.Name = ipaArgs.Name
			ipa.BundleID = ipaArgs.BundleID
			ipa.IsInterim = cast.ToInt(ipaArgs.IsInterim)
			util.PanicIf(h.ipaDAO.Update(ctx, ipa))
		}
		for _, version := range ipaArgs.Versions {
			ipaVersion, err := h.ipaVersionDAO.GetByIpaIDVersion(ctx, ipaID, version.Version)
			if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
				panic(err)
			}
			if ipaVersion == nil {
				util.PanicIf(h.ipaVersionDAO.Insert(ctx, &models.IpaVersion{
					IpaID:     ipaID,
					Version:   version.Version,
					TokenPath: version.Token,
				}))
				continue
			}
			ipaVersion.TokenPath = version.Token
			util.PanicIf(h.ipaVersionDAO.Update(ctx, ipaVersion))
		}
		/// 删除最久的一个 ipa version, 保证库里永远只存入三个 ipa
		versions := ipaVersionMap[ipaID]
		if len(versions) > 3 {
			version := versions[len(versions)-1]
			util.PanicIf(h.ipaVersionDAO.Delete(ctx, version.ID))
			util.PanicIf(h.tencentCtl.DeleteFile(ctx, version.TokenPath))
		}
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
		models.SearchRecordV2Where.IpaID.IN(ipaIDs),
		models.SearchRecordV2Where.CreatedAt.GTE(time.Date(0, 0, -3, 0, 0, 0, 0, time.Now().Location())),
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

type deleteIpaArgs struct {
	IpaID      string `json:"ipa_id" validate:"required"`
	IpaVersion string `json:"ipa_version"`
}

func (p *deleteIpaArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminIpaHandler) DeleteIpa(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)
	if _, ok := constant.OpsAuthMemberIDMap[loginID]; !ok {
		panic(errors.ErrMemberAccessDenied)
	}

	args := &deleteIpaArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	ipaID := cast.ToInt64(args.IpaID)
	if args.IpaVersion == "" {
		util.PanicIf(h.deleteIpa(ctx, ipaID))
	} else {
		util.PanicIf(h.deleteIpaVersion(ctx, ipaID, args.IpaVersion))
	}

	util.RenderJSON(w, "ok")
}

func (h *AdminIpaHandler) deleteIpa(ctx context.Context, ipaID int64) error {
	ipa, err := h.ipaDAO.Get(ctx, ipaID)
	if err != nil {
		return err
	}
	ivs, err := h.ipaVersionDAO.GetIpaVersionSliceByIpaID(ctx, ipaID)
	if err != nil {
		return err
	}

	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	err = h.ipaDAO.Delete(ctx, ipa.ID)
	if err != nil {
		return err
	}

	for _, iv := range ivs {
		err = h.ipaVersionDAO.Delete(ctx, iv.ID)
		if err != nil {
			return err
		}
		err = h.tencentCtl.DeleteFile(ctx, iv.TokenPath)
		if err != nil {
			return err
		}
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	return nil
}

func (h *AdminIpaHandler) deleteIpaVersion(ctx context.Context, ipaID int64, ipaVersion string) error {
	iv, err := h.ipaVersionDAO.GetByIpaIDVersion(ctx, ipaID, ipaVersion)
	if err != nil {
		return err
	}
	err = h.ipaVersionDAO.Delete(ctx, iv.ID)
	if err != nil {
		return err
	}
	return h.tencentCtl.DeleteFile(ctx, iv.TokenPath)
}
