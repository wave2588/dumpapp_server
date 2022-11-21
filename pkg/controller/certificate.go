package controller

import (
	"context"

	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
)

type CertificateController interface {
	CreateCer(ctx context.Context, UDID, regionPool string) *CertificateResponse
	CheckCerIsActive(ctx context.Context, certificateID int64) (bool, error)
	GetBalance(ctx context.Context) (*CertificateBalance, error)
}

type CertificateResponse struct {
	P12Data             string                     `json:"p12_data"`
	MobileProvisionData string                     `json:"mobile_provision_data"`
	Source              enum.CertificateSource     `json:"source"`
	BizExt              datatype.CertificateBizExt `json:"biz_ext"`

	/// 生成错误这里会有错误原因
	ErrorMessage *string `json:"error_message"`
}

type CertificateBalance struct {
	Count int64 `json:"count"`
}
