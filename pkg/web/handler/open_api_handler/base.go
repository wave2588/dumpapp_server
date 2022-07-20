package open_api_handler

import (
	"context"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller/impl"
)

func mustGetLoginID(ctx context.Context, r *http.Request) int64 {
	token := r.Header.Get("x-dumpapp-token")
	memberID, err := impl.DefaultMemberIDEncryptionController.GetMemberIDByCode(ctx, token)
	util.PanicIf(err)
	return memberID
}
