package handler

import (
	dao2 "dumpapp_server/pkg/dao"
	impl4 "dumpapp_server/pkg/dao/impl"
	"net/http"
)

type MemberDownloadRecordHandler struct {
	accountDAO dao2.AccountDAO
}

func NewMemberDownloadRecordHandler() *MemberDownloadRecordHandler {
	return &MemberDownloadRecordHandler{
		accountDAO: impl4.DefaultAccountDAO,
	}
}

func (h *MemberDownloadRecordHandler) GetSelfDownloadRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	account := GetAccountByLoginID(ctx, loginID)

}
