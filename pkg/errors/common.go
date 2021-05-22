package errors

var (
	HttpBadRequestError           = NewDefaultAPIError(400, 400, "BadRequestError", "请求错误")
	HttpForbiddenError            = NewDefaultAPIError(403, 106, "ForbiddenError", "没有对应的权限")
	PermissionDeniedError         = NewDefaultAPIError(403, 403, "PermissionDeniedError", "没有对应的权限")
	HttpNotFoundError             = NewDefaultAPIError(404, 4041, "NotFoundError", "资源不存在")
	ErrMethodNotAllowed           = NewDefaultAPIError(405, 405, "MethodNotAllowed", "HTTP 方法不支持")
	HttpResourceNotAvailableError = NewDefaultAPIError(410, 107, "ResourceNotAvailable", "资源不可用")
	HttpUnprocessableError        = NewDefaultAPIError(422, 4000, "UnprocessableError", "请求不能被处理")
	HttpServerError               = NewDefaultAPIError(500, 5000, "ServerInternalError", "服务器内部异常")
)

type DetailError struct {
	Code         int          `json:"code"`
	Name         string       `json:"name"`
	Message      string       `json:"message"`
	Data         *interface{} `json:"data,omitempty"`
	DebugMessage *string      `json:"debug_message,omitempty"`
}

type APIError struct {
	DetailError    DetailError `json:"error"`
	httpStatusCode int
}

func (e *APIError) Error() string {
	return e.DetailError.Message
}

func (e *APIError) HttpStatus() int {
	return e.httpStatusCode
}

func NewErrorFromTemplate(err *APIError, msg string) *APIError {
	return NewDefaultAPIError(err.httpStatusCode, err.DetailError.Code, err.DetailError.Name, msg)
}

func NewDefaultAPIError(statusCode int, errorCode int, kind, message string) *APIError {
	return &APIError{
		DetailError: DetailError{
			Code:         errorCode,
			Name:         kind,
			Message:      message,
			Data:         nil,
			DebugMessage: nil,
		},
		httpStatusCode: statusCode,
	}
}

func NewAPIError(statusCode int, errorCode int, kind, message string, data *interface{}, debugMessage *string) *APIError {
	return &APIError{
		DetailError: DetailError{
			Code:         errorCode,
			Name:         kind,
			Message:      message,
			Data:         data,
			DebugMessage: debugMessage,
		},
		httpStatusCode: statusCode,
	}
}

func NewMalformRequestException(message string, data *interface{}, debugMessage *string) *APIError {
	if message == "" {
		message = "请求错误"
	}
	return NewAPIError(400, 400, "MalformRequestException", message, data, debugMessage)
}

func NewUnproccessableException(message string, data *interface{}, debugMessage *string) *APIError {
	if message == "" {
		message = "请求不能被处理"
	}
	return NewAPIError(422, 4000, "UnproccessableException", message, data, debugMessage)
}

func NewResourceNotFoundException(message string, data *interface{}, debugMessage *string) *APIError {
	if message == "" {
		message = "资源不存在"
	}
	return NewAPIError(404, 4041, "ResourceNotFoundException", message, data, debugMessage)
}

func NewSpamRequestException(message string, data *interface{}, debugMessage *string) *APIError {
	if message == "" {
		message = "检测到异常情况,请求失败"
	}
	return NewAPIError(403, 106, "SpamRequestException", message, data, debugMessage)
}

func NewResourceNotAvailableException(message string, data *interface{}, debugMessage *string) *APIError {
	if message == "" {
		message = "资源不可用"
	}
	return NewAPIError(410, 107, "ResourceNotAvailableException", message, data, debugMessage)
}
