package handler

import (
	"fmt"
	"net/http"

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
	ExpenseID       string `json:"expense_id" validate:"required"` /// 客户端生成的消费 ID
	IpaName         string `json:"ipa_name" validate:"required"`
	IpaBundleID     string `json:"ipa_bundle_id" validate:"required"`
	IpaFileToken    string `json:"ipa_file_token" validate:"required"`
	IpaVersion      string `json:"ipa_version" validate:"required"`
	IpaSize         int64  `json:"ipa_size" validate:"required"`
	CertificateName string `json:"certificate_name" validate:"required"` /// 证书名称
}

func (args *postSignIpaArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *MemberSignIpaHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	args := &postSignIpaArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	/// 检查 ExpenseID 是否存在
	res, err := h.memberSignDAO.BatchGetByExpenseID(ctx, []string{args.ExpenseID})
	util.PanicIf(err)
	if _, ok := res[args.ExpenseID]; ok {
		util.PanicIf(errors.UnproccessableError("expense_id 已存在"))
	}

	bucket := config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket

	/// 获取签名后的 ipa 文件地址
	ipaURL, err := h.lingshulianCtl.GetURL(ctx, bucket, args.IpaFileToken)
	util.PanicIf(err)

	/// 开始上传
	plistToken := fmt.Sprintf("%d.plist", util2.MustGenerateID(ctx))
	util.PanicIf(h.fileCtl.PutFileToLocal(ctx, h.fileCtl.GetPlistFolderPath(ctx), plistToken, []byte(fmt.Sprintf(constant.MemberSignIpaPlistConfig, ipaURL, args.IpaBundleID, args.IpaName))))

	signIpaID := util2.MustGenerateID(ctx)
	util.PanicIf(h.memberSignDAO.Insert(ctx, &models.MemberSignIpa{
		ID:                signIpaID,
		ExpenseID:         args.ExpenseID,
		MemberID:          loginID,
		IsDelete:          false,
		IpaFileToken:      args.IpaFileToken,
		IpaPlistFileToken: plistToken,
		BizExt: datatype.MemberSignIpaBizExt{
			IpaName:         args.IpaName,
			IpaBundleID:     args.IpaBundleID,
			IpaVersion:      args.IpaVersion,
			IpaSize:         args.IpaSize,
			CertificateName: args.CertificateName,
		},
	}))

	data := render.NewMemberSignIpaRender([]int64{signIpaID}, loginID, render.MemberSignIpaDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, data[signIpaID])
}

func (h *MemberSignIpaHandler) GetSelfSignIpaList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var (
		loginID = mustGetLoginID(ctx)
		offset  = GetIntArgument(r, "offset", 0)
		limit   = GetIntArgument(r, "limit", 10)
	)

	filter := []qm.QueryMod{
		models.MemberSignIpaWhere.MemberID.EQ(loginID),
		models.MemberSignIpaWhere.IsDelete.EQ(false),
	}
	ids, err := h.memberSignDAO.ListIDs(ctx, offset, limit, filter, nil)
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

	fields := render.DefaultMemberSignIpaFields
	fields = append(fields, convertIncludes(args.IncludeFields)...)
	dataMap := render.NewMemberSignIpaRender(
		[]int64{id},
		0, []render.MemberSignIpaOption{
			render.MemberSignIpaIncludes(fields),
		}...,
	).RenderMap(ctx)

	data, ok := dataMap[id]
	if !ok {
		util.PanicIf(errors.ErrNotFound)
	}
	if data.IsDelete {
		util.PanicIf(errors.ErrNotFound)
	}

	/// 检查签名用户是否有足够的下载次数
	if len(args.IncludeFields) != 0 {
		util.PanicIf(h.dispenseCountCtl.Check(ctx, data.Meta.MemberID, 1))
	}

	util.RenderJSON(w, data)
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

	/// 删除 cos 记录
	_ = h.lingshulianCtl.Delete(ctx, config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket, data.Meta.IpaPlistFileToken)
	_ = h.lingshulianCtl.Delete(ctx, config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket, data.Meta.IpaFileToken)

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}
