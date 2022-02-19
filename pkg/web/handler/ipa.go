package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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

type IpaHandler struct {
	ipaDAO            dao.IpaDAO
	ipaVersionDAO     dao.IpaVersionDAO
	searchRecordV2DAO dao.SearchRecordV2DAO

	memberDownloadCtl controller.MemberDownloadController
	alterWebCtl       controller2.AlterWebController

	adminDumpOrderCtl controller.AdminDumpOrderController
}

func NewIpaHandler() *IpaHandler {
	return &IpaHandler{
		ipaDAO:            impl.DefaultIpaDAO,
		ipaVersionDAO:     impl.DefaultIpaVersionDAO,
		searchRecordV2DAO: impl.DefaultSearchRecordV2DAO,

		memberDownloadCtl: impl2.DefaultMemberDownloadController,
		alterWebCtl:       impl3.DefaultAlterWebController,

		adminDumpOrderCtl: impl2.DefaultAdminDumpOrderController,
	}
}

type getIpaArgs struct {
	Name         string       `form:"name" validate:"required"`
	BundleID     string       `form:"bundle_id"`
	Version      string       `form:"version"`
	AppStoreLink string       `from:"appstorelink"`
	IpaType      enum.IpaType `json:"ipa_type"`
}

func (p *getIpaArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if !p.IpaType.IsAIpaType() {
		p.IpaType = enum.IpaTypeNormal
	}
	return nil
}

func (h *IpaHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	ipaID := cast.ToInt64(util.URLParam(r, "ipa_id"))

	args := getIpaArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	/// 记录用户获取的记录
	util.PanicIf(h.searchRecordV2DAO.Insert(ctx, &models.SearchRecordV2{
		MemberID: loginID,
		IpaID:    ipaID,
		Name:     args.Name,
	}))

	ipa, err := h.ipaDAO.Get(ctx, ipaID)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}
	/// 如果找到了, 正常返回结构即可, 子页面会判断是否有下载次数
	if ipa != nil {
		data := render.NewIpaRender([]int64{ipaID}, loginID, []enum.IpaType{args.IpaType}, render.IpaDefaultRenderFields...).RenderMap(ctx)
		if len(data[ipaID].Versions) != 0 {
			util.RenderJSON(w, data[ipaID])
			return
		}
	}

	/// 判断是否有下载次数
	err = h.memberDownloadCtl.CheckPayCount(ctx, loginID, 9)
	util.PanicIf(err)

	/// 如果有下载次数, 并且库里没有这个 ipa 则去发送邮件
	util.RenderJSON(w, map[string]bool{
		"send_email": true,
	})
}

type getLatestVersionArgs struct {
	Name     string `form:"name" validate:"required"`
	BundleID string `form:"bundle_id"`
	Version  string `form:"version"`
}

func (p *getLatestVersionArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *IpaHandler) GetLatestVersion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	ipaID := cast.ToInt64(util.URLParam(r, "ipa_id"))

	args := getLatestVersionArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	/// 判断是否有下载次数
	err := h.memberDownloadCtl.CheckPayCount(ctx, loginID, 9)
	util.PanicIf(err)

	/// 库内没有找到对应的砸壳信息，需要发送推送给负责人进行砸壳。
	h.alterWebCtl.SendDumpOrderMsg(ctx, loginID, ipaID, args.BundleID, args.Name, args.Version)
}

func (h *IpaHandler) GetAllVersion(w http.ResponseWriter, r *http.Request) {
	ipaID := cast.ToInt64(util.URLParam(r, "ipa_id"))
	country := cast.ToString(util.URLParam(r, "country"))
	endpoint := fmt.Sprintf("https://tools.lancely.tech/api/apple/appVersion/%s/%d", country, ipaID)
	body, err := util2.HttpRequest("GET", endpoint, map[string]string{}, map[string]string{})
	util.PanicIf(err)
	var result interface{}
	util.PanicIf(json.Unmarshal(body, &result))
	util.RenderJSON(w, result)
}

type getRankingArgs struct {
	StartAt int64 `form:"start_at"`
	EndAt   int64 `form:"end_at"`
}

func (args *getRankingArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

type searchRanking struct {
	IpaID int64  `json:"ipa_id,string"`
	Name  string `json:"name"`
}

func (h *IpaHandler) GetRanking(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := getRankingArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	filter := make([]qm.QueryMod, 0)
	if args.StartAt != 0 {
		filter = append(filter, models.SearchRecordV2Where.CreatedAt.GTE(cast.ToTime(args.StartAt)))
	}
	if args.EndAt != 0 {
		filter = append(filter, models.SearchRecordV2Where.CreatedAt.LTE(cast.ToTime(args.EndAt)))
	}

	data, err := h.searchRecordV2DAO.GetOrderBySearchCount(ctx, 0, 20, filter)
	util.PanicIf(err)

	result := make([]*searchRanking, 0)
	for _, datum := range data {
		result = append(result, &searchRanking{
			IpaID: datum.IpaID,
			Name:  datum.Name,
		})
	}

	util.RenderJSON(w, util.ListOutput{
		Paging: nil,
		Data:   result,
	})
}
