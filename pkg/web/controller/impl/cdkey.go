package impl

import (
	"context"

	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
)

type CdkeyWebController struct {
	memberPayCountCtl controller.MemberPayCountController
}

var DefaultCdkeyWebController *CdkeyWebController

func init() {
	DefaultCdkeyWebController = NewCdkeyWebController()
}

func NewCdkeyWebController() *CdkeyWebController {
	return &CdkeyWebController{
		memberPayCountCtl: impl2.DefaultMemberPayCountController,
	}
}

func (c *CdkeyWebController) Create(ctx context.Context, memberID int64) (int64, error) {
	return 0, nil
}
