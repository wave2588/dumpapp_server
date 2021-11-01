package http

import "context"

type CertificateServer interface {
	CreateCer(ctx context.Context, udid string) (*CreateCerResponse, error)
}

type CreateCerResponse struct {
	IsSuccess    bool        `json:"IsSuccess"`
	Data         *Data       `json:"data"`
	ErrorCode    int         `json:"ErrorCode"`
	ErrorMessage interface{} `json:"ErrorMessage"`
}

type Data struct {
	P12FileDate             string `json:"p12_file_date"`
	MobileProvisionFileData string `json:"mobile_provision_file_data"`
	UdidBatchNo             string `json:"udid_batch_no"`
	CerAppleid              string `json:"cer_appleid"`
}
