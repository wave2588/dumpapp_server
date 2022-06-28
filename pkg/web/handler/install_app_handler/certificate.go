package install_app_handler

import (
	"context"
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	impl3 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	controller2 "dumpapp_server/pkg/web/controller"
	impl2 "dumpapp_server/pkg/web/controller/impl"
	"dumpapp_server/pkg/web/render/install_app_render"
	"github.com/go-playground/validator/v10"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CertificateHandler struct {
	certificateV2Controller  controller.CertificateController
	alterWebCtl              controller2.AlterWebController
	certificateWebCtl        controller2.CertificateWebController
	installAppCertificateDAO dao.InstallAppCertificateDAO
	installAppCDKeyDAO       dao.InstallAppCdkeyDAO
}

func NewCertificateHandler() *CertificateHandler {
	return &CertificateHandler{
		certificateV2Controller:  impl.DefaultCertificateV2Controller,
		alterWebCtl:              impl2.DefaultAlterWebController,
		certificateWebCtl:        impl2.DefaultCertificateWebController,
		installAppCertificateDAO: impl3.DefaultInstallAppCertificateDAO,
		installAppCDKeyDAO:       impl3.DefaultInstallAppCdkeyDAO,
	}
}

type postCertificate struct {
	UDID  string `json:"udid" validate:"required"`
	CDKey string `json:"cdkey" validate:"required"`
}

func (p *postCertificate) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *CertificateHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &postCertificate{}
	util.PanicIf(util.JSONArgs(r, args))

	cdkeyMap, err := h.installAppCDKeyDAO.BatchGetByOutID(ctx, []string{args.CDKey})
	util.PanicIf(err)

	cdkey, ok := cdkeyMap[args.CDKey]
	if !ok {
		util.PanicIf(errors.ErrInstallAppCDKeyNotFound)
	}
	/// 兑换码如果已经被使用了，则需要直接返回数据即可。
	if cdkey.Status == enum.InstallAppCDKeyStatusUsed {
		h.renderData(ctx, w, cdkey.ID)
		return
	}
	/// 兑换码被删除了则报错处理
	if cdkey.Status == enum.InstallAppCDKeyStatusAdminDelete {
		util.PanicIf(errors.ErrInstallAppCDKeyAdminDelete)
	}

	/// 请求整数接口
	response := h.certificateV2Controller.CreateCer(ctx, args.UDID, "1")
	if response.ErrorMessage != nil {
		/// 创建失败推送
		h.alterWebCtl.SendInstallAppCreateCertificateFailMsg(ctx, args.CDKey, args.UDID, *response.ErrorMessage)
		util.PanicIf(errors.ErrCreateCertificateFail)
	}
	if response.BizExt == nil {
		util.PanicIf(errors.ErrCreateCertificateFail)
	}

	p12FileData := response.P12Data
	mpFileData := response.MobileProvisionData
	/// p12 文件修改内容
	modifiedP12FileData, err := h.certificateWebCtl.GetModifiedCertificateData(ctx, p12FileData, response.BizExt.OriginalP12Password, response.BizExt.NewP12Password)
	util.PanicIf(err)

	/// 计算证书 md5
	p12FileMd5 := util2.StringMd5(p12FileData)
	mpFileMd5 := util2.StringMd5(mpFileData)

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	cerID := util2.MustGenerateID(ctx)

	cdkey.Status = enum.InstallAppCDKeyStatusUsed
	cdkey.CertificateID = cerID
	util.PanicIf(h.installAppCDKeyDAO.Update(ctx, cdkey))

	util.PanicIf(h.installAppCertificateDAO.Insert(ctx, &models.InstallAppCertificate{
		ID:                         cerID,
		Udid:                       args.UDID,
		P12FileData:                p12FileData,
		P12FileDataMD5:             p12FileMd5,
		ModifiedP12FileDate:        modifiedP12FileData,
		MobileProvisionFileData:    mpFileData,
		MobileProvisionFileDataMD5: mpFileMd5,
		Source:                     response.Source,
		BizExt:                     response.BizExt.String(),
	}))

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	h.renderData(ctx, w, cdkey.ID)
}

func (h *CertificateHandler) renderData(ctx context.Context, w http.ResponseWriter, cdkeyID int64) {
	data := install_app_render.NewCDKEYRender([]int64{cdkeyID}, 0, install_app_render.CDKeyDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, data[cdkeyID])
}

func (h *CertificateHandler) GetByUDID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		udid   = util.URLParam(r, "udid")
		offset = GetIntArgument(r, "offset", 0)
		limit  = GetIntArgument(r, "limit", 10)
	)

	filter := []qm.QueryMod{
		models.InstallAppCertificateWhere.Udid.EQ(udid),
	}
	ids, err := h.installAppCertificateDAO.ListIDs(ctx, offset, limit, filter, nil)
	util.PanicIf(err)
	count, err := h.installAppCertificateDAO.Count(ctx, filter)
	util.PanicIf(err)

	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(count), offset, limit),
		Data:   install_app_render.NewCertificateRender(ids, 0, install_app_render.CertificateDefaultRenderFields...).RenderSlice(ctx),
	})
}
