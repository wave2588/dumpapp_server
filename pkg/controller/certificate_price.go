package controller

import "context"

type CertificatePriceController interface {
	BatchGetPrices(ctx context.Context, memberIDs []int64) (map[int64][]*CertificatePriceInfo, error)
	GetPrices(ctx context.Context, memberID int64) ([]*CertificatePriceInfo, error)
}

type CertificatePriceInfo struct {
	ID          int64  `json:"id,string"`
	Price       int64  `json:"price"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
