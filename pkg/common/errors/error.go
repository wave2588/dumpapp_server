package errors

import (
	"fmt"
)

type SError struct {
	code    int
	name    string
	message string
}

func (e *SError) Error() string {
	return fmt.Sprintf("<%s> code=%d, Message=%s", e.name, e.code, e.message)
}

func (e *SError) Code() int {
	return e.code
}

func (e *SError) Name() string {
	return e.name
}

func (e *SError) Message() string {
	return e.message
}

func (e *SError) SetMessage(msg string) error {
	e.message += "\n" + msg
	return e
}

func NewError(code int, name, message string) *SError {
	return &SError{
		code:    code,
		name:    name,
		message: message,
	}
}

var (
	ErrAccessDenied    = NewError(403, "AccessDenied", "没有权限进行操作")
	ErrNotFound        = NewError(404, "NotFound", "记录不存在")
	ErrDuplicatedEntry = NewError(409, "DuplicatedEntry", "记录已存在")
	ErrUnprocessable   = NewError(422, "Unprocessable", "无法处理")
	ErrTooManyRequests = NewError(509, "TooManyRequest", "操作太快，请稍后再试")
)

var ErrNotFoundApp = NewError(404, "NotFoundApp", "未找到 app")
