package controller

import "context"

type AlterWebController interface {
	/// 推送自定义信息
	SendCustomMsg(ctx context.Context, token, content string)
	/// 砸壳订单推送
	SendDumpOrderMsg(ctx context.Context, loginID, ipaID int64, bundleID, ipaName, version string)
	/// feedback 推送
	SendFeedbackMsg(ctx context.Context, loginID int64, content string)
	/// 创建证书失败推送
	SendCreateCertificateFailMsg(ctx context.Context, loginID, deviceID int64, errorMessage string)
	/// 安装 app 创建证书失败推送
	SendInstallAppCreateCertificateFailMsg(ctx context.Context, cdkey, udid string, errorMessage string)
	/// 证书创建成功
	SendBeganCreateCertificateMsg(ctx context.Context, loginID int64, udid string, isReplenish bool)
	SendCreateCertificateSuccessMsg(ctx context.Context, loginID, deviceID, cerID int64) // 废弃了
	SendCreateCertificateSuccessMsgV2(ctx context.Context, loginID, deviceID, cerID int64, isReplenish bool)

	///  注册用户监控
	SendAccountMsg(ctx context.Context)

	/// 绑定设备流程 log
	SendDeviceLog(ctx context.Context, title string, memberID int64, values map[string]string)
}
