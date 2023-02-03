package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/controller"
	impl2 "dumpapp_server/pkg/web/controller/impl"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AdminCertificateHandler struct {
	accountDAO        dao.AccountDAO
	memberDeviceDAO   dao.MemberDeviceDAO
	certificateDAO    dao.CertificateV2DAO
	certificateWebCtl controller.CertificateWebController
}

func NewAdminCertificateHandler() *AdminCertificateHandler {
	return &AdminCertificateHandler{
		accountDAO:        impl.DefaultAccountDAO,
		memberDeviceDAO:   impl.DefaultMemberDeviceDAO,
		certificateDAO:    impl.DefaultCertificateV2DAO,
		certificateWebCtl: impl2.DefaultCertificateWebController,
	}
}

type replenishCertificateArgs struct {
	Email string `json:"email" validate:"required"`
	UDID  string `json:"udid" validate:"required"`
}

func (p *replenishCertificateArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if !util2.CheckUDIDValid(p.UDID) {
		return errors.UnproccessableError(fmt.Sprintf("无效的 UDID: %s", p.UDID))
	}
	return nil
}

func (h *AdminCertificateHandler) Replenish(w http.ResponseWriter, r *http.Request) {
	util.PanicIf(errors.UnproccessableError("该接口已下线"))

	//ctx := r.Context()
	//
	//args := &replenishCertificateArgs{}
	//util.PanicIf(util.JSONArgs(r, args))
	//
	//accountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{args.Email})
	//util.PanicIf(err)
	//
	//account, ok := accountMap[args.Email]
	//if !ok {
	//	util.PanicIf(errors.UnproccessableError("邮箱未找到"))
	//}
	//
	//devices, err := h.memberDeviceDAO.GetByMemberIDAndUDIDs(ctx, account.ID, []string{args.UDID})
	//util.PanicIf(err)
	//
	//if len(devices) == 0 {
	//	util.PanicIf(errors.UnproccessableError(fmt.Sprintf("当前账号下没有此 UDID: %s", args.UDID)))
	//}
	//
	//device := devices[0]
	//cerIDs, err := impl.DefaultCertificateV2DAO.ListIDs(ctx, 0, 100, []qm.QueryMod{
	//	models.CertificateV2Where.DeviceID.EQ(device.ID),
	//}, nil)
	//util.PanicIf(err)
	//
	//if len(cerIDs) == 0 {
	//	util.PanicIf(errors.UnproccessableError("该账号下的 UDID 没有购买过证书 UDID"))
	//}
	//
	//cerMap := render.NewCertificateRender(cerIDs, 0, render.CertificateDefaultRenderFields...).RenderMap(ctx)
	//
	//var cer *render.Certificate
	//for _, cerID := range cerIDs {
	//	cer = cerMap[cerID]
	//	if !cer.IsReplenish {
	//		break
	//	}
	//}
	//if cer == nil {
	//	util.PanicIf(errors.UnproccessableError("未找到有效证书"))
	//}
	//
	//// 检查证书是否有效
	//if cer.P12IsActive {
	//	util.PanicIf(errors.UnproccessableError("证书有效，无法候补。"))
	//}
	//
	//// 0 说明是老版本证书, 需要管理员校验
	//if cer.Level == 0 {
	//	util.PanicIf(errors.UnproccessableError("当前证书无法候补，请联系管理员。"))
	//}
	//
	//now := time.Now()
	//if cer.ReplenishExpireAt <= now.Unix() {
	//	switch cer.Level {
	//	case 1:
	//		util.PanicIf(errors.UnproccessableError("已超过 7 天候补时间，无法候补。"))
	//	case 2:
	//		util.PanicIf(errors.UnproccessableError("已超过 180 天候补时间，无法候补。"))
	//	case 3:
	//		util.PanicIf(errors.UnproccessableError("已超过 365 天候补时间，无法候补。"))
	//	}
	//}
	//
	//_, err = h.certificateWebCtl.PayCertificate(ctx, account.ID, args.UDID, "售后证书", constant.CertificateIDL1, true, "")
	//util.PanicIf(err)
	//
	//util.RenderJSON(w, DefaultSuccessBody(ctx))
}

type getCertificateArgs struct {
	UDID string `form:"udid"`
}

func (args *getCertificateArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminCertificateHandler) GetCertificate(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		loginID = mustGetLoginID(ctx)
	)
	args := getCertificateArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	deviceIDs, err := h.memberDeviceDAO.ListIDs(ctx, 0, 100, []qm.QueryMod{
		models.MemberDeviceWhere.Udid.EQ(args.UDID),
	}, []string{})
	util.PanicIf(err)

	fmt.Println(deviceIDs)
	cerIDMap, err := h.certificateDAO.ListIDsByDeviceIDs(ctx, deviceIDs)
	util.PanicIf(err)

	certificateIDs := make([]int64, 0)
	for _, ids := range cerIDMap {
		certificateIDs = append(certificateIDs, ids...)
	}

	data := render.NewCertificateRender(certificateIDs, loginID, render.CertificateDefaultRenderFields...).RenderSlice(ctx)

	util.RenderJSON(w, util.ListOutput{
		Paging: nil,
		Data:   data,
	})
}
