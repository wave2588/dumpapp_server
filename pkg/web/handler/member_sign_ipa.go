package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MemberSignIpaHandler struct {
	lingshulianCtl   controller.LingshulianController
	fileCtl          controller.FileController
	dispenseCountCtl controller.DispenseCountController
	memberSignDAO    dao.MemberSignIpaDAO
}

func NewMemberSignIpaHandler() *MemberSignIpaHandler {
	return &MemberSignIpaHandler{
		lingshulianCtl:   impl.DefaultLingshulianController,
		fileCtl:          impl.DefaultFileController,
		dispenseCountCtl: impl.DefaultDispenseCountController,
		memberSignDAO:    impl2.DefaultMemberSignIpaDAO,
	}
}

type postSignIpaArgs struct {
	IpaName         string `json:"ipa_name" validate:"required"`
	IpaBundleID     string `json:"ipa_bundle_id" validate:"required"`
	IpaFileToken    string `json:"ipa_file_token" validate:"required"`
	IpaVersion      string `json:"ipa_version" validate:"required"`
	IpaSize         int64  `json:"ipa_size" validate:"required"`
	CertificateName string `json:"certificate_name" validate:"required"` /// 证书名称
	DispenseCount   *int64 `json:"dispense_count"`
	IsDumpapp       bool   `json:"is_dumpapp"`
	Note            string `json:"note"`
	AppTimeLockID   int64  `json:"app_time_lock_id,string"`
}

func (args *postSignIpaArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if util2.StringCount(args.Note) > 50 {
		return errors.UnproccessableError("备注字数请小于 50 字")
	}
	return nil
}

func (h *MemberSignIpaHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	args := &postSignIpaArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	bucket := config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket
	/// 获取签名后的 ipa 文件地址
	ipaURL, err := h.lingshulianCtl.GetURL(ctx, bucket, args.IpaFileToken)
	util.PanicIf(err)
	/// 开始上传
	plistToken := fmt.Sprintf("%d.plist", util2.MustGenerateID(ctx))
	util.PanicIf(h.fileCtl.PutFileToLocal(ctx, h.fileCtl.GetPlistFolderPath(ctx), plistToken, []byte(fmt.Sprintf(constant.MemberSignIpaPlistConfig, ipaURL, args.IpaBundleID, args.IpaName))))

	// 如果是 dumpapp 客户端更新，则只返回 plist_url
	if args.IsDumpapp && args.IpaName == "DumpApp" && args.IpaBundleID == "com.dumpapp.ipa" {
		util.RenderJSON(w, map[string]string{
			"plist_url": h.fileCtl.GetPlistFileURL(ctx, plistToken),
		})
		return
	}

	signIpaID := util2.MustGenerateID(ctx)
	bizExt := datatype.MemberSignIpaBizExt{
		IpaName:         args.IpaName,
		IpaBundleID:     args.IpaBundleID,
		IpaVersion:      args.IpaVersion,
		IpaSize:         args.IpaSize,
		CertificateName: args.CertificateName,
		AppTimeLockID:   args.AppTimeLockID,
	}
	if args.DispenseCount != nil {
		bizExt.DispenseCount = *args.DispenseCount
	}
	util.PanicIf(h.memberSignDAO.Insert(ctx, &models.MemberSignIpa{
		ID:                signIpaID,
		ExpenseID:         util2.MustGenerateUUID(),
		MemberID:          loginID,
		IsDelete:          false,
		IpaFileToken:      args.IpaFileToken,
		IpaPlistFileToken: plistToken,
		Note:              args.Note,
		BizExt:            bizExt,
	}))

	data := render.NewMemberSignIpaRender([]int64{signIpaID}, loginID, render.MemberSignIpaDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, data[signIpaID])
}

type putSignIpaArgs struct {
	DispenseCount *int64 `json:"dispense_count" validate:"required"`
	Note          string `json:"note"`
}

func (args *putSignIpaArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if util2.StringCount(args.Note) > 50 {
		return errors.UnproccessableError("备注字数请小于 50 字")
	}
	return nil
}

func (h *MemberSignIpaHandler) Put(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		id      = cast.ToInt64(util.URLParam(r, "id"))
		loginID = mustGetLoginID(ctx)
	)

	args := &putSignIpaArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	signIpaMap, err := h.memberSignDAO.BatchGet(ctx, []int64{id})
	util.PanicIf(err)

	signIpa, ok := signIpaMap[id]
	if !ok {
		util.PanicIf(errors.ErrNotFound)
	}

	if signIpa.IsDelete {
		util.PanicIf(errors.UnproccessableError("已删除的 ipa 无法修改"))
	}
	if signIpa.MemberID != loginID {
		util.PanicIf(errors.ErrMemberAccessDenied)
	}

	if args.DispenseCount != nil {
		signIpa.BizExt.DispenseCount = *args.DispenseCount
	}
	signIpa.Note = args.Note
	util.PanicIf(h.memberSignDAO.Update(ctx, signIpa))

	data := render.NewMemberSignIpaRender([]int64{id}, loginID, render.MemberSignIpaDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, data[id])
}

type getMemberSignIpaListArgs struct {
	StartAt int64  `form:"start_at"`
	EndAt   int64  `form:"end_at"`
	Keyword string `form:"keyword"`
}

func (args *getMemberSignIpaListArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *MemberSignIpaHandler) GetSelfSignIpaList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var (
		loginID = mustGetLoginID(ctx)
		offset  = GetIntArgument(r, "offset", 0)
		limit   = GetIntArgument(r, "limit", 10)
	)

	args := getMemberSignIpaListArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	filter := []qm.QueryMod{
		models.MemberSignIpaWhere.MemberID.EQ(loginID),
		models.MemberSignIpaWhere.IsDelete.EQ(false),
	}
	if args.StartAt != 0 {
		filter = append(filter, models.MemberSignIpaWhere.CreatedAt.GTE(time.Unix(args.StartAt, 0)))
	}
	if args.EndAt != 0 {
		filter = append(filter, models.MemberSignIpaWhere.CreatedAt.LTE(time.Unix(args.EndAt, 0)))
	}
	if args.Keyword != "" {
		filter = append(filter, qm.Where("note like ?", fmt.Sprintf(`%s%%`, args.Keyword)))
	}

	ids, err := h.memberSignDAO.ListIDs(ctx, offset, limit, filter, []string{"updated_at desc"})
	util.PanicIf(err)

	totalCount, err := h.memberSignDAO.Count(ctx, filter)
	util.PanicIf(err)

	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(totalCount), offset, limit),
		Data:   render.NewMemberSignIpaRender(ids, loginID, render.MemberSignIpaDefaultRenderFields...).RenderSlice(ctx),
	})
}

type getMemberSignIpaArgs struct {
	IncludeFields datatype.IncludeFields `form:"include"`
}

func (args *getMemberSignIpaArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *MemberSignIpaHandler) Get(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = cast.ToInt64(util.URLParam(r, "id"))
	)

	args := getMemberSignIpaArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	data, err := h.getMemberSignIpaData(ctx, id, args.IncludeFields)
	util.PanicIf(err)

	util.RenderJSON(w, data)
}

type getMemberSignIpaV2Args struct {
	IncludeFields datatype.IncludeFields `form:"include"`
	IDs           Int64StringSlice       `form:"ids"`
}

func (args *getMemberSignIpaV2Args) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *MemberSignIpaHandler) GetV2(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := getMemberSignIpaV2Args{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	fmt.Println(args.IDs)

	result := make([]*render.MemberSignIpa, 0)
	for _, id := range args.IDs {
		data, err := h.getMemberSignIpaData(ctx, id, args.IncludeFields)
		if err != nil {
			continue
		}
		result = append(result, data)
	}

	fields := render.DefaultMemberSignIpaFields
	fields = append(fields, convertIncludes(args.IncludeFields)...)

	util.RenderJSON(w, util.ListOutput{
		Paging: nil,
		Data:   result,
	})
}

// Deprecated
func (h *MemberSignIpaHandler) GetByExpenseID(w http.ResponseWriter, r *http.Request) {
	util.PanicIf(errors.UnproccessableError("不需要处理"))
}

func (h *MemberSignIpaHandler) getMemberSignIpaData(ctx context.Context, memberSignIpaID int64, includeFields datatype.IncludeFields) (*render.MemberSignIpa, error) {
	fields := render.DefaultMemberSignIpaFields
	fields = append(fields, convertIncludes(includeFields)...)
	dataMap := render.NewMemberSignIpaRender(
		[]int64{memberSignIpaID},
		0, []render.MemberSignIpaOption{
			render.MemberSignIpaIncludes(fields),
		}...,
	).RenderMap(ctx)

	data, ok := dataMap[memberSignIpaID]
	if !ok || data.IsDelete {
		return nil, errors.ErrNotFound
	}

	/// 检查签名用户是否有足够的下载次数
	if len(includeFields) != 0 {
		dCount := h.dispenseCountCtl.CalculateMemberSignIpaDispenseCount(ctx, data.Meta.BizExt.IpaSize)
		err := h.dispenseCountCtl.Check(ctx, data.Meta.MemberID, dCount)
		if err != nil {
			return nil, err
		}

		/// 如果分发次数已经 > 用户设置分发次数, 则不允许再分发
		if !data.Dispense.IsValid {
			if err = errors.UnproccessableError("分发次数已达到上限"); err != nil {
				return nil, err
			}
		}
	}
	return data, nil
}

func convertIncludes(includeFields datatype.IncludeFields) []string {
	supportFieldMap := map[string]string{
		"plist_url": "PlistURL",
	}
	includes := make([]string, 0)
	for _, field := range includeFields {
		if f, ok := supportFieldMap[field]; ok {
			includes = append(includes, f)
		}
	}
	return includes
}

func (h *MemberSignIpaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		loginID = mustGetLoginID(ctx)
		id      = cast.ToInt64(util.URLParam(r, "id"))
	)

	dataMap := render.NewMemberSignIpaRender([]int64{id}, loginID, render.MemberSignIpaDefaultRenderFields...).RenderMap(ctx)
	data, ok := dataMap[id]
	if !ok {
		util.PanicIf(errors.ErrNotFound)
	}
	if data.Meta.MemberID != loginID {
		util.PanicIf(errors.ErrMemberAccessDenied)
	}

	meta := data.Meta
	meta.IsDelete = true

	/// 标记删除记录
	util.PanicIf(h.memberSignDAO.Update(ctx, meta))

	/// 删除本地文件
	plistPath := fmt.Sprintf("%s/%s", h.fileCtl.GetPlistFolderPath(ctx), data.Meta.IpaPlistFileToken)
	_ = h.fileCtl.DeleteFile(ctx, plistPath)

	/// 删除 cos 记录
	_ = h.lingshulianCtl.Delete(ctx, config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket, data.Meta.IpaPlistFileToken)
	_ = h.lingshulianCtl.Delete(ctx, config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket, data.Meta.IpaFileToken)

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}
