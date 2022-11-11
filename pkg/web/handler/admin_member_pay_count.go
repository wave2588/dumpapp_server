package handler

import (
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	controller2 "dumpapp_server/pkg/web/controller"
	impl3 "dumpapp_server/pkg/web/controller/impl"
	"github.com/go-playground/validator/v10"
)

type AdminMemberPayCountHandler struct {
	accountDAO        dao.AccountDAO
	memberPayCountCtl controller.MemberPayCountController
	alertWebCtl       controller2.AlterWebController
}

func NewAdminMemberPayCountHandler() *AdminMemberPayCountHandler {
	return &AdminMemberPayCountHandler{
		accountDAO:        impl.DefaultAccountDAO,
		memberPayCountCtl: impl2.DefaultMemberPayCountController,
		alertWebCtl:       impl3.DefaultAlterWebController,
	}
}

type addDownloadNumber struct {
	Email  string `json:"email" validate:"required"`
	Number int64  `json:"number" validate:"required"`
}

func (p *addDownloadNumber) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminMemberPayCountHandler) AddNumber(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	args := &addDownloadNumber{}
	util.PanicIf(util.JSONArgs(r, args))

	accountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{args.Email})
	util.PanicIf(err)

	account, ok := accountMap[args.Email]
	if !ok {
		util.PanicIf(errors.ErrNotFoundMember)
		return
	}

	util.PanicIf(h.memberPayCountCtl.AddCount(ctx, account.ID, args.Number, enum.MemberPayCountSourceAdminPresented, datatype.MemberPayCountRecordBizExt{
		ObjectID:   0,
		ObjectType: datatype.MemberPayCountRecordBizExtObjectTypeNone,
	}))

	// 加个推送
	adminAccountMap, err := h.accountDAO.BatchGet(ctx, []int64{loginID})
	util.PanicIf(err)
	adminAccount := adminAccountMap[loginID]
	titleString := "<font color=\"warning\">管理员添加 D 币</font>\n>"
	countString := fmt.Sprintf("count：<font color=\"comment\">%d</font>\n", args.Number)
	receiveEmailString := fmt.Sprintf("用户邮箱：<font color=\"comment\">%s</font>\n", args.Email)
	adminEmailString := fmt.Sprintf("管理员邮箱：<font color=\"comment\">%s</font>\n", adminAccount.Email)
	timeStr := fmt.Sprintf("操作时间：<font color=\"comment\">%s</font>", time.Now().Format("2006-01-02 15:04:05"))
	h.alertWebCtl.SendCustomMsg(ctx, "32df4de7-524c-4d0c-94cd-c8d7e0709fb4", titleString+countString+receiveEmailString+adminEmailString+timeStr)
}

type deleteDownloadNumber struct {
	Email  string `json:"email" validate:"required"`
	Number int64  `json:"number" validate:"required"`
}

func (p *deleteDownloadNumber) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminMemberPayCountHandler) DeleteNumber(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := mustGetLoginID(ctx)

	args := &deleteDownloadNumber{}
	util.PanicIf(util.JSONArgs(r, args))

	accountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{args.Email})
	util.PanicIf(err)

	account, ok := accountMap[args.Email]
	if !ok {
		util.PanicIf(errors.ErrNotFoundMember)
		return
	}

	util.PanicIf(h.memberPayCountCtl.CheckPayCount(ctx, account.ID, args.Number))

	util.PanicIf(h.memberPayCountCtl.DeductPayCount(ctx, account.ID, args.Number, enum.MemberPayCountStatusAdminDelete, enum.MemberPayCountUseAdminDelete, datatype.MemberPayCountRecordBizExt{
		ObjectID:   0,
		ObjectType: datatype.MemberPayCountRecordBizExtObjectTypeNone,
	}))

	// 加个推送
	adminAccountMap, err := h.accountDAO.BatchGet(ctx, []int64{loginID})
	util.PanicIf(err)
	adminAccount := adminAccountMap[loginID]
	titleString := "<font color=\"warning\">管理员删除 D 币</font>\n>"
	countString := fmt.Sprintf("count：<font color=\"comment\">%d</font>\n", args.Number)
	receiveEmailString := fmt.Sprintf("用户邮箱：<font color=\"comment\">%s</font>\n", args.Email)
	adminEmailString := fmt.Sprintf("管理员邮箱:：<font color=\"comment\">%s</font>\n", adminAccount.Email)
	timeStr := fmt.Sprintf("操作时间：<font color=\"comment\">%s</font>", time.Now().Format("2006-01-02 15:04:05"))
	h.alertWebCtl.SendCustomMsg(ctx, "32df4de7-524c-4d0c-94cd-c8d7e0709fb4", titleString+countString+receiveEmailString+adminEmailString+timeStr)
}
