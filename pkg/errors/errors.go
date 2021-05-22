package errors

var (
	ErrNotAuthorized  = NewDefaultAPIError(401, 10000, "NotAuthorized", "登陆才可以继续操作")
	ErrInvalidTicket  = NewDefaultAPIError(401, 10001, "InvalidTicket", "无效的用户身份")
	ErrNotFoundMember = NewDefaultAPIError(404, 10002, "NotFoundMember", "未找到用户")
	ErrUpgradeVip     = NewDefaultAPIError(401, 10003, "UpgradeVip", "请升级 vip")

	ErrNotFoundIpaVersion = NewDefaultAPIError(404, 20003, "NotFoundIpaVersion", "未找到对应的 ipa 版本")
	ErrNotFoundApp        = NewDefaultAPIError(404, 20001, "NotFoundApp", "未找到 app")
)

func UnproccessableError(msg string) *APIError {
	return NewDefaultAPIError(422, 4000, "Unproccessable", msg)
}
