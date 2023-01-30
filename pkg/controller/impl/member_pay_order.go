package impl

import "dumpapp_server/pkg/controller"

type MemberPayOrderController struct{}

var DefaultMemberPayOrderController *MemberPayOrderController

func init() {
	DefaultMemberPayOrderController = NewMemberPayOrderController()
}

func NewMemberPayOrderController() *MemberPayOrderController {
	return &MemberPayOrderController{}
}

func (c *MemberPayOrderController) GetPayCampaignDescription() string {
	return "充 500 送 15 ，充 1500 送 100，充 3000 送 300 ，充 5000 送 800。"
}

func (c *MemberPayOrderController) GetPayCampaignRule() *controller.PayCampaign {
	return &controller.PayCampaign{
		Items: []*controller.PayCampaignItem{
			{
				PayCount:        500,
				PayForFreeCount: 15,
			},
			{
				PayCount:        1000,
				PayForFreeCount: 70,
			},
			{
				PayCount:        2000,
				PayForFreeCount: 260,
			},
			{
				PayCount:        5000,
				PayForFreeCount: 1290,
			},
		},
	}
}
