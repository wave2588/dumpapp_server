package handler

import (
	"fmt"
	"net/http"

	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/middleware"
	util2 "dumpapp_server/pkg/util"
	pkgErr "github.com/pkg/errors"
)

type MemberInviteHandler struct {
	memberInviteCodeDAO dao.MemberInviteCodeDAO
}

func NewMemberInviteHandler() *MemberInviteHandler {
	return &MemberInviteHandler{
		memberInviteCodeDAO: impl.DefaultMemberInviteCodeDAO,
	}
}

func (h *MemberInviteHandler) PostInviteURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	inviteCode, err := h.memberInviteCodeDAO.GetByMemberID(ctx, loginID)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}

	inviteCodeString := ""

	/// 如果邀请码已经存在, 则直接取出即可
	if inviteCode != nil {
		inviteCodeString = inviteCode.Code
	}

	/// 如果没有邀请码则生成邀请码
	if inviteCode == nil {
		inviteCodeString = util2.MustGenerateInviteCode(ctx, 6)
		util.PanicIf(h.memberInviteCodeDAO.Insert(ctx, &models.MemberInviteCode{
			MemberID: loginID,
			Code:     inviteCodeString,
		}))
	}
	util.RenderJSON(w, map[string]string{
		"url": fmt.Sprintf("https://www.dumpapp.com/register?invite_code=%s", inviteCodeString),
	})
}
