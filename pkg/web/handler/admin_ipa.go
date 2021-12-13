package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
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
	"dumpapp_server/pkg/web/render"
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

	appleCtl          controller.AppleController
	emailWebCtl       controller2.EmailWebController
	tencentCtl        controller.TencentController
	adminDumpOrderCtl controller.AdminDumpOrderController
}

func NewAdminIpaHandler() *AdminIpaHandler {
	return &AdminIpaHandler{
		accountDAO:        impl.DefaultAccountDAO,
		ipaDAO:            impl.DefaultIpaDAO,
		ipaVersionDAO:     impl.DefaultIpaVersionDAO,
		searchRecordV2DAO: impl.DefaultSearchRecordV2DAO,

		appleCtl:          impl2.DefaultAppleController,
		emailWebCtl:       impl3.DefaultEmailWebController,
		tencentCtl:        impl2.DefaultTencentController,
		adminDumpOrderCtl: impl2.DefaultAdminDumpOrderController,
	}
}

func (h *AdminIpaHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	offset := GetIntArgument(r, "offset", 0)
	limit := GetIntArgument(r, "limit", 10)

	ids, err := h.ipaDAO.ListIDs(ctx, offset, limit, nil, nil)
	util.PanicIf(err)

	totalCount, err := h.ipaDAO.Count(ctx, nil)
	util.PanicIf(err)

	ipa := render.NewIpaRender(ids, loginID, render.IpaAdminRenderFields...).RenderSlice(ctx)
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(totalCount), offset, limit),
		Data:   ipa,
	})
}

type createIpaArgs struct {
	Ipas        []*ipaArgs `json:"ipas" validate:"required"`
	IsSendEmail bool       `json:"is_send_email"`
}

type ipaArgs struct {
	IpaID    string     `json:"ipa_id" validate:"required"`
	Name     string     `json:"name" validate:"required"`
	BundleID string     `json:"bundle_id" validate:"required"`
	Versions []*Version `json:"versions"`
}

type Version struct {
	Version     string       `json:"version" validate:"required"`
	Token       string       `json:"token" validate:"required"`
	IpaType     enum.IpaType `json:"ipa_type" validate:"required"`
	DescribeURL *string      `json:"describe_url"`
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

	for _, ipaArgs := range args.Ipas {
		ipaID := cast.ToInt64(ipaArgs.IpaID)
		ipa := ipaMap[ipaID]
		if ipa == nil {
			util.PanicIf(h.ipaDAO.Insert(ctx, &models.Ipa{
				ID:       ipaID,
				Name:     ipaArgs.Name,
				BundleID: ipaArgs.BundleID,
				Type:     enum.IpaTypeNormal, /// todo: 要删掉
			}))
		} else {
			ipa.Name = ipaArgs.Name
			ipa.BundleID = ipaArgs.BundleID
			ipa.Type = enum.IpaTypeNormal /// todo: 要删掉
			util.PanicIf(h.ipaDAO.Update(ctx, ipa))
		}
		for _, version := range ipaArgs.Versions {
			/// 找出 ipa_id ipa_type version 相同的数据，全部删掉重新上传
			ipaVersions, err := h.ipaVersionDAO.GetByIpaIDAndIpaTypeAndVersion(ctx, ipaID, version.IpaType, version.Version)
			if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
				panic(err)
			}
			for _, ipaVersion := range ipaVersions {
				util.PanicIf(h.ipaVersionDAO.Delete(ctx, ipaVersion.ID))
			}

			/// 删除 dump order 记录
			util.PanicIf(h.adminDumpOrderCtl.Progressed(ctx, loginID, ipaID, version.Version))

			ipaVersionBizExt := &constant.IpaVersionBizExt{DescribeURL: version.DescribeURL}
			util.PanicIf(h.ipaVersionDAO.Insert(ctx, &models.IpaVersion{
				IpaID:     ipaID,
				Version:   version.Version,
				TokenPath: version.Token,
				IpaType:   version.IpaType,
				BizExt:    ipaVersionBizExt.String(),
			}))
		}
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	if args.IsSendEmail {
		util.PanicIf(h.sendEmail(ctx, ipaArgsMap))
	}

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

type batchDeleteIpaArgs struct {
	IpaIDs                []string `json:"ipa_ids" validate:"required"`
	IsRetainLatestVersion bool     `json:"is_retain_latest_version"` /// 是否保留最新版本
}

func (p *batchDeleteIpaArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminIpaHandler) BatchDeleteIpa(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)
	if _, ok := constant.OpsAuthMemberIDMap[loginID]; !ok {
		panic(errors.ErrMemberAccessDenied)
	}

	args := &batchDeleteIpaArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	ipaIDs := make([]int64, 0)
	for _, id := range args.IpaIDs {
		ipaIDs = append(ipaIDs, cast.ToInt64(id))
	}

	if args.IsRetainLatestVersion {
		util.PanicIf(h.batchDeleteByRetainLatestVersion(ctx, ipaIDs))
	} else {
		util.PanicIf(h.batchDeleteAll(ctx, ipaIDs))
	}
}

func (h *AdminIpaHandler) batchDeleteByRetainLatestVersion(ctx context.Context, ipaIDs []int64) error {
	ipaMap := render.NewIpaRender(ipaIDs, 0, render.IpaDefaultRenderFields...).RenderMap(ctx)

	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	for _, ipa := range ipaMap {
		for idx, version := range ipa.Versions {
			if idx == 0 {
				continue
			}
			err := h.ipaVersionDAO.Delete(ctx, version.ID)
			if err != nil {
				return err
			}
		}
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	return nil
}

func (h *AdminIpaHandler) batchDeleteAll(ctx context.Context, ipaIDs []int64) error {
	ipaMap := render.NewIpaRender(ipaIDs, 0, render.IpaDefaultRenderFields...).RenderMap(ctx)

	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	for _, ipa := range ipaMap {
		err := h.ipaDAO.Delete(ctx, ipa.ID)
		if err != nil {
			return err
		}
		for _, version := range ipa.Versions {
			err = h.ipaVersionDAO.Delete(ctx, version.ID)
			if err != nil {
				return err
			}
		}
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	return nil
}

type deleteIpaArgs struct {
	IpaID                 string       `json:"ipa_id" validate:"required"`
	IpaType               enum.IpaType `json:"ipa_type" validate:"required"`
	IpaVersion            string       `json:"ipa_version"`              /// 指定删除某个版本
	IsRetainLatestVersion bool         `json:"is_retain_latest_version"` /// 是否保留最新版本
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

	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	ipaID := cast.ToInt64(args.IpaID)
	if args.IpaVersion == "" {
		/// 保留最新版本, 删除其他的
		if args.IsRetainLatestVersion {
			util.PanicIf(h.deleteIpaRetainLatestVersion(ctx, ipaID, args.IpaType))
		} else {
			/// 删除操作
			util.PanicIf(h.deleteIpa(ctx, ipaID))
		}
	} else {
		util.PanicIf(h.deleteIpaVersion(ctx, ipaID, args.IpaType, args.IpaVersion))
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	util.RenderJSON(w, "ok")
}

func (h *AdminIpaHandler) deleteIpaRetainLatestVersion(ctx context.Context, ipaID int64, ipaType enum.IpaType) error {
	ivs, err := h.ipaVersionDAO.GetByIpaIDAndIpaType(ctx, ipaID, ipaType)
	if err != nil {
		return err
	}
	for idx, iv := range ivs {
		if idx == 0 {
			continue
		}
		err = h.ipaVersionDAO.Delete(ctx, iv.ID)
		if err != nil {
			return err
		}
		err = h.tencentCtl.DeleteFile(ctx, iv.TokenPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *AdminIpaHandler) deleteIpa(ctx context.Context, ipaID int64) error {
	ivs, err := h.ipaVersionDAO.GetByIpaID(ctx, ipaID)
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
	return nil
}

func (h *AdminIpaHandler) deleteIpaVersion(ctx context.Context, ipaID int64, ipaType enum.IpaType, ipaVersion string) error {
	ivs, err := h.ipaVersionDAO.GetByIpaIDAndIpaTypeAndVersion(ctx, ipaID, ipaType, ipaVersion)
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
	return nil
}
