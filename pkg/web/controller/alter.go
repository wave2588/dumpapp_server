package controller

import "context"

type AlterWebController interface {
	/// 砸壳订单推送
	SendDumpOrderMsg(ctx context.Context, loginID, ipaID int64, bundleID, ipaName, version string)
	/// feedback 推送
	SendFeedbackMsg(ctx context.Context, loginID int64, content string)
	/// 创建证书失败推送
	SendCreateCertificateFailMsg(ctx context.Context, loginID, deviceID int64, errorMessage string)
	/// 证书创建成功
	SendCreateCertificateSuccessMsg(ctx context.Context, loginID, deviceID, cerID int64)

	///  注册用户监控
	SendAccountMsg(ctx context.Context)

	/// 绑定设备流程 log
	SendDeviceLog(ctx context.Context, title string, memberID int64, values map[string]string)
}
