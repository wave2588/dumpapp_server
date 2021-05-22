package errors

var (
	ErrNotAuthorized = NewDefaultAPIError(401, 10000, "NotAuthorized", "登陆才可以继续操作")
	ErrInvalidTicket = NewDefaultAPIError(401, 10001, "InvalidTicket", "无效的用户身份")

	ErrNotFoundApp = NewDefaultAPIError(404, 20000, "NotFoundApp", "未找到 app")
)

func UnproccessableError(msg string) *APIError {
	return NewDefaultAPIError(422, 4000, "Unproccessable", msg)
}
