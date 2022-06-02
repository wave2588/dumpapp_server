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
	ErrMemberInviterTooMuch     = NewDefaultAPIError(403, 10013, "ErrMemberInviterTooMuch", "此邀请码今日邀请过多，请稍后重试。")
	ErrPhoneRefusedRegister     = NewDefaultAPIError(401, 10014, "PhoneRefusedRegister", "该手机号不允许注册。")
	ErrReLogin                  = NewDefaultAPIError(401, 10015, "ReLogin", "身份过期，请重新登录")

	ErrNotFoundApp        = NewDefaultAPIError(404, 20001, "NotFoundApp", "未找到 app")
	ErrNotFoundIpa        = NewDefaultAPIError(404, 20002, "NotFoundIpa", "未找到 ipa")
	ErrNotFoundIpaVersion = NewDefaultAPIError(404, 20003, "NotFoundIpaVersion", "未找到对应的 ipa 版本")

	/// 支付相关的错误
	ErrNotPayCount           = NewDefaultAPIError(403, 30001, "NotDownloadNumber", "D 币不足，请充值 D 币。")
	ErrCreateCertificateFail = NewDefaultAPIError(403, 30003, "CreateCertificateFail", "生成证书失败")

	/// 证书相关错误
	ErrNotFoundCertificate         = NewDefaultAPIError(404, 40000, "NotFoundCertificate", "未找到对应的证书")
	ErrCertificateInvalid          = NewDefaultAPIError(404, 40001, "CertificateInvalid", "证书已失效")
	ErrCertificateUnavailByAccount = NewDefaultAPIError(404, 40001, "CertificateUnavail", "当前账号不能使用此证书")

	/// 业务 dao 层错误
	ErrRedisFail       = NewDefaultAPIError(500, 50000, "RedisFail", "redis 发生错误")
	ErrMemberBlacklist = NewDefaultAPIError(403, 50001, "MemberBlacklist", "该账户已被拉黑, 请稍后重试。")

	/// 签名相关
	ErrNotFoundIpaSign          = NewDefaultAPIError(404, 60000, "NotFoundIpaSign", "未找到对应的签名任务")
	ErrIpaSignStatusProcessing  = NewDefaultAPIError(401, 60001, "IpaSignStatusProcessing", "签名任务进行中")
	ErrIpaSignStatusFail        = NewDefaultAPIError(401, 60002, "IpaSignStatusFail", "签名任务失败")
	ErrIpaSignStatusUnprocessed = NewDefaultAPIError(401, 60003, "IpaSignStatusFail", "签名任务未开始，请稍等。")

	/// 每日免费相关
	ErrDailyFreeNone   = NewDefaultAPIError(403, 70001, "ErrDailyFreeNone", "每日免费次数已用完")
	ErrDailyFreeUnique = NewDefaultAPIError(403, 70002, "ErrDailyFreeUnique", "每个人每天只有一次机会")

	/// 设备相关的错误
	ErrDeviceNotFound = NewDefaultAPIError(404, 80001, "ErrDailyFreeNone", "未找到对应的设备")

	///
	ErrHttpFail = NewDefaultAPIError(403, 403, "HttpFail", "http 请求失败")
	ErrNotFound = NewDefaultAPIError(404, 404, "NotFound", "记录未找到")
)

func UnproccessableError(msg string) *APIError {
	return NewDefaultAPIError(422, 4000, "Unproccessable", msg)
}
