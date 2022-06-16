package controller

import (
	"context"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
)

type CertificateController interface {
	CreateCer(ctx context.Context, UDID, regionPool string) *CertificateResponse
	CheckCerIsActive(ctx context.Context, certificateID int64) (bool, error)
}

type CertificateResponse struct {
	P12Data             string                      `json:"p12_data"`
	MobileProvisionData string                      `json:"mobile_provision_data"`
	Source              enum.CertificateSource      `json:"source"`
	BizExt              *constant.CertificateBizExt `json:"biz_ext"`

	/// 生成错误这里会有错误原因
	ErrorMessage *string `json:"error_message"`
}

///// 创建证书
//CreateCer(ctx context.Context, udid, regionPool string) (*CreateCerResponse, error)
///// 检测p12证书是否有效
//CheckP12File(ctx context.Context, p12FileData, p12Password string) (*CheckCerResponse, error)
///// 检测证书是否有效
//CheckCerByUDIDBatchNo(ctx context.Context, udidBatchNo string) (*CheckCerResponse, error)
