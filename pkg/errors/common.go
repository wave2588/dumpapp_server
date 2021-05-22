package errors

func UnproccessableError(msg string) *APIError {
	return NewDefaultAPIError(422, 4000, "Unproccessable", msg)
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

type DetailError struct {
	Code         int          `json:"code"`
	Name         string       `json:"name"`
	Message      string       `json:"message"`
	Data         *interface{} `json:"data,omitempty"`
	DebugMessage *string      `json:"debug_message,omitempty"`
}

// APIError represents an detailed error in ss standard API output format.
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

func NotAuthorizedError(msg string) *APIError {
	return NewDefaultAPIError(401, 401, "NotAuthorized", msg)
}

func InvalidTicketError(msg string) *APIError {
	return NewDefaultAPIError(401, 41300, "Member.InvalidTicket", msg)
}
