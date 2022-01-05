package errors

var (
	ErrNotAuthorized            = NewDefaultAPIError(401, 10000, "NotAuthorized", "登陆才可以继续操作")
	ErrInvalidTicket            = NewDefaultAPIError(401, 10001, "InvalidTicket", "无效的用户身份")
	ErrNotFoundMember           = NewDefaultAPIError(404, 10002, "NotFoundMember", "未找到用户")
	ErrUpgradeVip               = NewDefaultAPIError(401, 10003, "UpgradeVip", "请升级 vip")
	ErrMemberAccessDenied       = NewDefaultAPIError(403, 10004, "MemberAccessDenied", "没有权限进行操作")
	ErrAccountRegisteredByEmail = NewDefaultAPIError(403, 10005, "AccountRegisteredByEmail", "该邮箱已被注册")
	ErrAccountRegisteredByPhone = NewDefaultAPIError(403, 10006, "AccountRegisteredByPhone", "该手机号已被注册")
	ErrCaptchaIncorrectByEmail  = NewDefaultAPIError(403, 10007, "CaptchaIncorrectByEmail", "邮箱验证码错误")
	ErrCaptchaIncorrectByPhone  = NewDefaultAPIError(403, 10008, "CaptchaIncorrectByPhone", "手机验证码错误")
	ErrCaptchaRepeated          = NewDefaultAPIError(403, 10009, "CaptchaRepeated", "验证码重复发送")
	ErrMemberInviteCodeInvalid  = NewDefaultAPIError(403, 10010, "ErrMemberInviteCodeInvalid", "邀请码无效")
	ErrAccountUnusual           = NewDefaultAPIError(401, 10011, "AccountUnusual", "账户异常，请联系管理员。")
	ErrEmailRefusedRegister     = NewDefaultAPIError(401, 10012, "EmailRefusedRegister", "该邮箱不允许注册。")

	ErrNotFoundApp        = NewDefaultAPIError(404, 20001, "NotFoundApp", "未找到 app")
	ErrNotFoundIpa        = NewDefaultAPIError(404, 20002, "NotFoundIpa", "未找到 ipa")
	ErrNotFoundIpaVersion = NewDefaultAPIError(404, 20003, "NotFoundIpaVersion", "未找到对应的 ipa 版本")

	/// 支付相关的错误
	ErrNotDownloadNumber         = NewDefaultAPIError(403, 30001, "NotDownloadNumber", "没有下载次数")
	ErrDownloadNumberLessThanSix = NewDefaultAPIError(403, 30002, "DownloadNumberLessThanFive", "下载次数不足 6 次")
	ErrCreateCertificateFail     = NewDefaultAPIError(403, 30003, "CreateCertificateFail", "生成证书失败")

	/// 证书相关错误
	ErrNotFoundCertificate = NewDefaultAPIError(404, 40000, "NotFoundCertificate", "未找到对应的证书")
	ErrCertificateInvalid  = NewDefaultAPIError(404, 40001, "CertificateInvalid", "证书已失效")

	/// 签名相关
	ErrNotFoundIpaSign          = NewDefaultAPIError(404, 60000, "NotFoundIpaSign", "未找到对应的签名任务")
	ErrIpaSignStatusProcessing  = NewDefaultAPIError(401, 60001, "IpaSignStatusProcessing", "签名任务进行中")
	ErrIpaSignStatusFail        = NewDefaultAPIError(401, 60002, "IpaSignStatusFail", "签名任务失败")
	ErrIpaSignStatusUnprocessed = NewDefaultAPIError(401, 60003, "IpaSignStatusFail", "签名任务未开始，请稍等。")

	/// 业务 dao 层错误
	ErrRedisFail       = NewDefaultAPIError(500, 50000, "RedisFail", "redis 发生错误")
	ErrMemberBlacklist = NewDefaultAPIError(403, 50001, "MemberBlacklist", "该账户已被拉黑, 请稍后重试。")

	///
	ErrHttpFail = NewDefaultAPIError(403, 403, "HttpFail", "http 请求失败")
	ErrNotFound = NewDefaultAPIError(404, 404, "NotFound", "记录未找到")
)

func UnproccessableError(msg string) *APIError {
	return NewDefaultAPIError(422, 4000, "Unproccessable", msg)
}
