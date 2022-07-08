package dao

import "context"

type IpaRankingDAO interface {
	SetIpaRankingData(ctx context.Context, data *IpaRanking) error
	GetIpaRankingData(ctx context.Context) (*IpaRanking, error)
	RemoveIpaRankingData(ctx context.Context) error
}

type IpaRanking struct {
	Data []interface{} `json:"data"`
}
