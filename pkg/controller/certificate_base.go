package controller

import (
	"context"
	"time"

	"dumpapp_server/pkg/dao/models"
)

type CertificateBaseController interface {
	// 查看证书是否有效
	CheckCertificateIsActive(ctx context.Context, ids []int64) (map[int64]bool, error)
	CheckCertificateIsActiveByModels(ctx context.Context, certificates models.CertificateV2Slice) (map[int64]bool, error)
	// 获取证书候补时间
	GetCertificateReplenishExpireAt(ctx context.Context, ids []int64) (map[int64]time.Time, error)
	GetCertificateReplenishExpireAtByModels(ctx context.Context, certificates models.CertificateV2Slice) (map[int64]time.Time, error)

	// 创建证书
	Create(ctx context.Context, UDID string) (cerBase *CerCreateResponse, alterMsg string, err error)
}

type CerCreateResponse struct {
	Response            *CertificateResponse
	ModifiedP12FileData string
}
