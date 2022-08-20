package handler

import (
	"fmt"
	"net/http"
	"strings"

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
	lingshulianCtl controller.LingshulianController
	memberSignDAO  dao.MemberSignIpaDAO
}

func NewMemberSignIpaHandler() *MemberSignIpaHandler {
	return &MemberSignIpaHandler{
		lingshulianCtl: impl.DefaultLingshulianController,
		memberSignDAO:  impl2.DefaultMemberSignIpaDAO,
	}
}

type postSignIpaArgs struct {
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

	/// 获取签名后的 ipa 文件地址
	ipaURL, err := h.lingshulianCtl.GetURL(ctx, config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket, args.IpaFileToken)
	util.PanicIf(err)

	plistToken := fmt.Sprintf("%d.plist", util2.MustGenerateID(ctx))
	/// 获取上传凭证
	putResp, err := h.lingshulianCtl.GetPutURL(ctx, config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket, plistToken)
	util.PanicIf(err)
	/// 开始上传
	util.PanicIf(h.lingshulianCtl.PutFile(ctx, putResp.URL, strings.NewReader(fmt.Sprintf(constant.MemberSignIpaPlistConfig, ipaURL, args.IpaBundleID, args.IpaName))))

	signIpaID := util2.MustGenerateID(ctx)
	util.PanicIf(h.memberSignDAO.Insert(ctx, &models.MemberSignIpa{
		ID:                signIpaID,
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

func (h *MemberSignIpaHandler) Get(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = cast.ToInt64(util.URLParam(r, "id"))
	)

	dataMap := render.NewMemberSignIpaRender([]int64{id}, 0, render.MemberSignIpaDefaultRenderFields...).RenderMap(ctx)
	data, ok := dataMap[id]
	if !ok {
		util.PanicIf(errors.ErrNotFound)
	}
	if data.IsDelete {
		util.PanicIf(errors.ErrNotFound)
	}
	util.RenderJSON(w, data)
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
