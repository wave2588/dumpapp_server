package http

import "context"

type CertificateServer interface {
	/// 创建证书
	CreateCer(ctx context.Context, udid string) (*CreateCerResponse, error)
	/// 检测证书是否有效
	CheckCer(ctx context.Context, p12FileData, p12Password string) (*CheckCerResponse, error)
}

type CreateCerResponse struct {
	IsSuccess    bool                   `json:"IsSuccess"`
	Data         *CreateCerResponseData `json:"Data"`
	ErrorCode    int                    `json:"ErrorCode"`
	ErrorMessage string                 `json:"ErrorMessage"`
}

type CreateCerResponseData struct {
	P12FileDate             string `json:"p12_file_date"`
	MobileProvisionFileData string `json:"mobile_provision_file_data"`
	UdidBatchNo             string `json:"udid_batch_no"`
	CerAppleid              string `json:"cer_appleid"`
}

type CheckCerResponse struct {
	IsSuccess    bool   `json:"IsSuccess"`
	Data         bool   `json:"Data"`
	ErrorCode    int    `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
}
