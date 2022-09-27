package open_api_handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	controller2 "dumpapp_server/pkg/web/controller"
	impl3 "dumpapp_server/pkg/web/controller/impl"
	"dumpapp_server/pkg/web/handler"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type OpenCertificateHandler struct {
	memberDeviceDAO   dao.MemberDeviceDAO
	certificateDAO    dao.CertificateV2DAO
	certificateWebCtl controller2.CertificateWebController
}

func NewOpenCertificateHandler() *OpenCertificateHandler {
	return &OpenCertificateHandler{
		memberDeviceDAO:   impl.DefaultMemberDeviceDAO,
		certificateDAO:    impl.DefaultCertificateV2DAO,
		certificateWebCtl: impl3.DefaultCertificateWebController,
	}
}

type postCertificateArgs struct {
	UDID               string  `json:"udid" validate:"required"`
	CertificatePriceID *string `json:"certificate_price_id"`
}

func (p *postCertificateArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if !util2.CheckUDIDValid(p.UDID) {
		return errors.UnproccessableError(fmt.Sprintf("无效的 UDID: %s", p.UDID))
	}
	return nil
}

func (h *OpenCertificateHandler) PostCertificate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx, r)

	args := &postCertificateArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	device, err := h.memberDeviceDAO.GetByMemberIDUdidSafe(ctx, loginID, args.UDID)
	util.PanicIf(err)

	/// 判断是否绑定了设备，如果没有绑定则进行绑定
	deviceID := util2.MustGenerateID(ctx)
	if device != nil {
		deviceID = device.ID
	} else {
		util.PanicIf(h.memberDeviceDAO.Insert(ctx, &models.MemberDevice{
			ID:       deviceID,
			MemberID: loginID,
			Udid:     args.UDID,
			BizExt:   datatype.MemberDeviceBizExt{},
		}))
	}

	/// 计算证书价格
	cerPrice := constant.CertificatePriceL1
	if args.CertificatePriceID != nil {
		switch *args.CertificatePriceID {
		case "1":
			cerPrice = constant.CertificatePriceL1
		case "2":
			cerPrice = constant.CertificatePriceL2
		case "3":
			cerPrice = constant.CertificatePriceL3
		default:
			util.PanicIf(errors.UnproccessableError("certificate_price_id 未识别"))
		}
	}
	/// 购买证书
	cerID, err := h.certificateWebCtl.PayCertificate(ctx, loginID, args.UDID, cerPrice, "")
	util.PanicIf(err)

	cerMap := render.NewCertificateRender([]int64{cerID}, loginID, render.CertificateDefaultRenderFields...).RenderMap(ctx)
	cer, ok := cerMap[cerID]
	if !ok {
		util.PanicIf(errors.ErrNotFoundCertificate)
	}
	util.RenderJSON(w, cer)
}

type getCertificateArgs struct {
	CertificateID string `form:"certificate_id" validate:"required"`
}

func (p *getCertificateArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *OpenCertificateHandler) GetCertificate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := getCertificateArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	certificateID := cast.ToInt64(args.CertificateID)

	loginID := mustGetLoginID(ctx, r)

	cerMap := render.NewCertificateRender([]int64{certificateID}, loginID, render.CertificateDefaultRenderFields...).RenderMap(ctx)
	cer, ok := cerMap[certificateID]
	if !ok {
		util.PanicIf(errors.ErrNotFoundCertificate)
	}
	util.RenderJSON(w, cer)
}

type getCertificateListArgs struct {
	UDID *string `form:"udid"`
}

func (p *getCertificateListArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.UDID != nil {
		if !util2.CheckUDIDValid(*p.UDID) {
			return errors.UnproccessableError(fmt.Sprintf("无效的 UDID: %s", p.UDID))
		}
	}
	return nil
}

func (h *OpenCertificateHandler) GetCertificateList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		offset  = handler.GetIntArgument(r, "offset", 0)
		limit   = handler.GetIntArgument(r, "limit", 10)
		loginID = mustGetLoginID(ctx, r)
	)

	args := getCertificateListArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	deviceIDs := make([]int64, 0)
	/// 如果没有给 udid, 则返回全量数据
	if args.UDID == nil {
		deviceMap, err := h.memberDeviceDAO.BatchGetByMemberIDs(ctx, []int64{loginID})
		util.PanicIf(err)
		for _, devices := range deviceMap {
			for _, device := range devices {
				deviceIDs = append(deviceIDs, device.ID)
			}
		}
	} else {
		devices, err := h.memberDeviceDAO.GetByMemberIDAndUDIDs(ctx, loginID, []string{*args.UDID})
		util.PanicIf(err)
		for _, device := range devices {
			deviceIDs = append(deviceIDs, device.ID)
		}
	}

	filter := []qm.QueryMod{
		models.CertificateV2Where.DeviceID.IN(deviceIDs),
	}
	cerIDs, err := h.certificateDAO.ListIDs(ctx, offset, limit, filter, nil)
	util.PanicIf(err)

	totalCount, err := h.certificateDAO.Count(ctx, filter)
	util.PanicIf(err)

	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(totalCount), offset, limit),
		Data:   render.NewCertificateRender(cerIDs, loginID, render.CertificateDefaultRenderFields...).RenderSlice(ctx),
	})
}

func (h *OpenCertificateHandler) GetCertificatePrice(w http.ResponseWriter, r *http.Request) {
	util.RenderJSON(w, util.ListOutput{
		Data: constant.GetCertificatePrices(),
	})
}
