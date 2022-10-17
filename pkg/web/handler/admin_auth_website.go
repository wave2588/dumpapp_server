package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
)

type AdminAuthWebsiteHandler struct {
	accountDAO          dao.AccountDAO
	adminAuthWebsiteDAO dao.AdminAuthWebsiteDAO
}

func NewAdminAuthWebsiteHandler() *AdminAuthWebsiteHandler {
	return &AdminAuthWebsiteHandler{
		accountDAO:          impl.DefaultAccountDAO,
		adminAuthWebsiteDAO: impl.DefaultAdminAuthWebsiteDAO,
	}
}

type authWebsiteArgs struct {
	Email  string `json:"email" validate:"required"`
	Domain string `json:"domain" validate:"required"`
}

func (p *authWebsiteArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminAuthWebsiteHandler) Auth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &authWebsiteArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	accountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{args.Email})
	util.PanicIf(err)

	account, ok := accountMap[args.Email]
	if !ok {
		util.PanicIf(errors.ErrNotFoundMember)
	}

	authWebsite, err := h.adminAuthWebsiteDAO.GetByDomainSafe(ctx, args.Domain)
	util.PanicIf(err)

	if authWebsite != nil {
		util.PanicIf(errors.UnproccessableError(fmt.Sprintf("%s 已存在", args.Domain)))
		return
	}

	util.PanicIf(h.adminAuthWebsiteDAO.Insert(ctx, &models.AdminAuthWebsite{
		MemberID: account.ID,
		Domain:   args.Domain,
		BizExt: datatype.AdminAuthWebsiteBizExt{
			IsOpen: true,
		},
	}))

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}

type unAuthWebsiteArgs struct {
	Domain string `json:"domain" validate:"required"`
}

func (p *unAuthWebsiteArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminAuthWebsiteHandler) UnAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &unAuthWebsiteArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	authWebsite, err := h.adminAuthWebsiteDAO.GetByDomainSafe(ctx, args.Domain)
	util.PanicIf(err)

	if authWebsite == nil {
		util.PanicIf(errors.UnproccessableError(fmt.Sprintf("%s 未找到", args.Domain)))
		return
	}

	util.PanicIf(h.adminAuthWebsiteDAO.Delete(ctx, authWebsite.ID))

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}

type authWebsiteItem struct {
	ID        int64          `json:"id,string"`
	Domain    string         `json:"domain"`
	IsOpen    bool           `json:"is_open"` /// 是否打开了
	Member    *render.Member `json:"member"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
}

func (h *AdminAuthWebsiteHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		offset = GetIntArgument(r, "offset", 0)
		limit  = GetIntArgument(r, "limit", 10)
	)

	totalCount, err := h.adminAuthWebsiteDAO.Count(ctx, nil)
	util.PanicIf(err)

	ids, err := h.adminAuthWebsiteDAO.ListIDs(ctx, offset, limit, nil, nil)
	util.PanicIf(err)

	authWebsiteMap, err := h.adminAuthWebsiteDAO.BatchGet(ctx, ids)
	util.PanicIf(err)

	memberIDs := make([]int64, 0)
	for _, id := range ids {
		aw, ok := authWebsiteMap[id]
		if !ok {
			continue
		}
		memberIDs = append(memberIDs, aw.MemberID)
	}
	memberMap := render.NewMemberRender(memberIDs, 0, render.MemberAdminRenderFields...).RenderMap(ctx)

	data := make([]*authWebsiteItem, 0)
	for _, id := range ids {
		aw, ok := authWebsiteMap[id]
		if !ok {
			continue
		}
		data = append(data, &authWebsiteItem{
			ID:        aw.ID,
			Domain:    aw.Domain,
			IsOpen:    aw.BizExt.IsOpen,
			Member:    memberMap[aw.MemberID],
			CreatedAt: aw.CreatedAt.Unix(),
			UpdatedAt: aw.UpdatedAt.Unix(),
		})
	}

	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(totalCount), offset, limit),
		Data:   data,
	})
}
