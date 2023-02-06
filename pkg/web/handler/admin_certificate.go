package handler

import (
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/util"
	controller2 "dumpapp_server/pkg/controller"
	impl3 "dumpapp_server/pkg/controller/impl"
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
	accountDAO          dao.AccountDAO
	memberDeviceDAO     dao.MemberDeviceDAO
	certificateDAO      dao.CertificateV2DAO
	certificateWebCtl   controller.CertificateWebController
	certificateV2WebCtl controller.CertificateV2WebController
	certificateBaseCtl  controller2.CertificateBaseController
}

func NewAdminCertificateHandler() *AdminCertificateHandler {
	return &AdminCertificateHandler{
		accountDAO:          impl.DefaultAccountDAO,
		memberDeviceDAO:     impl.DefaultMemberDeviceDAO,
		certificateDAO:      impl.DefaultCertificateV2DAO,
		certificateWebCtl:   impl2.DefaultCertificateWebController,
		certificateV2WebCtl: impl2.DefaultCertificateV2WebController,
		certificateBaseCtl:  impl3.DefaultCertificateBaseController,
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
	ctx := r.Context()

	args := &replenishCertificateArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	accountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{args.Email})
	util.PanicIf(err)

	account, ok := accountMap[args.Email]
	if !ok {
		util.PanicIf(errors.UnproccessableError("邮箱未找到"))
	}

	devices, err := h.memberDeviceDAO.GetByMemberIDAndUDIDs(ctx, account.ID, []string{args.UDID})
	util.PanicIf(err)

	if len(devices) == 0 {
		util.PanicIf(errors.UnproccessableError(fmt.Sprintf("当前账号下没有此 UDID: %s", args.UDID)))
	}

	device := devices[0]
	cerIDs, err := impl.DefaultCertificateV2DAO.ListIDs(ctx, 0, 100, []qm.QueryMod{
		models.CertificateV2Where.DeviceID.EQ(device.ID),
	}, nil)
	util.PanicIf(err)

	if len(cerIDs) == 0 {
		util.PanicIf(errors.UnproccessableError("该账号下的 UDID 没有购买过证书 UDID"))
	}
	cerID := cerIDs[0]

	cerMap, err := h.certificateBaseCtl.GetCertificateReplenishExpireAt(ctx, []int64{cerID})
	util.PanicIf(err)

	cerExpireAt, ok := cerMap[cerID]
	if !ok {
		util.PanicIf(errors.UnproccessableError("未获取到证书的过期时间"))
	}

	if cerExpireAt.Unix() < time.Now().Unix() {
		util.PanicIf(errors.UnproccessableError("已不在候补时间内"))
	}

	_, err = h.certificateV2WebCtl.AdminCreate(ctx, account.ID, args.UDID)
	util.PanicIf(err)

	util.RenderJSON(w, DefaultSuccessBody(ctx))
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
