package controller

type MemberPayOrderController interface {
	GetPayCampaignDescription() string
	GetPayCampaignRule() *PayCampaign
}

type PayCampaign struct {
	Items []*PayCampaignItem `json:"items"`
}

type PayCampaignItem struct {
	PayCount        int64 `json:"pay_count"`
	PayForFreeCount int64 `json:"pay_for_free_count"`
}
