package install_app_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller/install_app"
	"dumpapp_server/pkg/controller/install_app/impl"
	dao2 "dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render/install_app_render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CDKEYOrderHandler struct {
	aliPayCtl               install_app.ALiPayInstallAppController
	installAppCDKEYDAO      dao2.InstallAppCdkeyDAO
	installAppCDKEYOrderDAO dao2.InstallAppCdkeyOrderDAO
}

func NewCDKEYOrderHandler() *CDKEYOrderHandler {
	return &CDKEYOrderHandler{
		aliPayCtl:               impl.DefaultALiPayInstallAppController,
		installAppCDKEYDAO:      impl2.DefaultInstallAppCdkeyDAO,
		installAppCDKEYOrderDAO: impl2.DefaultInstallAppCdkeyOrderDAO,
	}
}

type getCDKEYOrderArgs struct {
	Number     int64  `form:"number" validate:"required"`
	ContactWay string `form:"contact_way" validate:"required"`
}

func (args *getCDKEYOrderArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if args.Number <= 0 {
		return errors.UnproccessableError("number > 0")
	}
	return nil
}

func (h *CDKEYOrderHandler) GetOrderURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := getCDKEYOrderArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	orderID, payURL, err := h.aliPayCtl.GetPayURLByInstallApp(ctx, args.Number, args.ContactWay)
	util.PanicIf(err)

	res := map[string]interface{}{
		"order_id": cast.ToString(orderID),
		"open_url": payURL,
	}
	util.RenderJSON(w, res)
}

func (h *CDKEYOrderHandler) GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orderID := cast.ToInt64(util.URLParam(r, "order_id"))
	offset := GetIntArgument(r, "offset", 0)
	limit := GetIntArgument(r, "limit", 10)

	filter := []qm.QueryMod{
		models.InstallAppCdkeyWhere.OrderID.EQ(orderID),
	}
	ids, err := h.installAppCDKEYDAO.ListIDs(ctx, offset, limit, filter, nil)
	util.PanicIf(err)

	count, err := h.installAppCDKEYDAO.Count(ctx, filter)
	util.PanicIf(err)

	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(count), offset, limit),
		Data:   install_app_render.NewCDKEYRender(ids, 0, install_app_render.CDKeyDefaultRenderFields...).RenderSlice(ctx),
	})
}

type getOrderByContactWatResp struct {
	ContactWat string                      `json:"contact_wat"`
	CDKeys     []*install_app_render.CDKEY `json:"cd_keys"`
}

func (h *CDKEYOrderHandler) GetOrderByContactWay(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contactWay := util.URLParam(r, "contact_way")

	var orderID int64

	offset := 0
	limit := 100
	hasNext := true
	filter := []qm.QueryMod{
		models.InstallAppCdkeyOrderWhere.Status.EQ(enum.MemberPayOrderStatusPaid),
	}
	for hasNext {
		if orderID != 0 {
			break
		}

		ids, err := h.installAppCDKEYOrderDAO.ListIDs(ctx, offset, limit, filter, nil)
		util.PanicIf(err)

		offset += len(ids)
		hasNext = len(ids) == limit

		orderMap, err := h.installAppCDKEYOrderDAO.BatchGet(ctx, ids)
		util.PanicIf(err)

		for _, order := range orderMap {
			var bizExt constant.InstallAppCDKEYOrderBizExt
			util.PanicIf(json.Unmarshal([]byte(order.BizExt), &bizExt))
			if bizExt.ContactWay == contactWay {
				orderID = order.ID
				break
			}
		}
	}

	if orderID == 0 {
		util.PanicIf(errors.UnproccessableError("未找到订单"))
	}
	resp, err := h.installAppCDKEYDAO.BatchGetByOrderIDs(ctx, []int64{orderID})
	util.PanicIf(err)

	cdkeyIDs := make([]int64, 0)
	for _, appCdkeys := range resp {
		for _, cdkey := range appCdkeys {
			cdkeyIDs = append(cdkeyIDs, cdkey.ID)
		}
	}
	util.RenderJSON(w, &getOrderByContactWatResp{
		ContactWat: contactWay,
		CDKeys:     install_app_render.NewCDKEYRender(cdkeyIDs, 0, install_app_render.CDKeyDefaultRenderFields...).RenderSlice(ctx),
	})
}
