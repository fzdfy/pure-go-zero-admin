package errorx

import (
	"errors"
	"net/http"
	"time"
)

const defaultCode = 1001

var (
	ErrNotFound = errors.New("data not find")
	//ErrorUserNotFound         = errors.New("用户不存在")
	//ErrorNoRequiredParameters = errors.New("必要参数不能为空")
	//ErrorUserOperation        = errors.New("用户正在操作中，请稍后重试")
)

type CodeError struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
}

type CodeErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Common errors code definition.
var message = map[int]string{
	http.StatusUnauthorized:        "Access token is invalid",
	http.StatusForbidden:           "Request forbidden",
	http.StatusNotFound:            "Not found",
	http.StatusProxyAuthRequired:   "No access permission",
	http.StatusConflict:            "Duplicate Submission",
	http.StatusPreconditionFailed:  "Request parameter missing or invalid",
	http.StatusUnprocessableEntity: "Invalid status",
	http.StatusInternalServerError: "Internal server error",
	http.StatusGatewayTimeout:      "Request timeout",
}

// Message returns a message based on given code.
func Message(code int) (string, bool) {
	v, ok := message[code]

	return v, ok
}

func NewError(code int) error {
	return &CodeError{Code: code}
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Message: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Message
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    e.Code,
		Message: e.Message,
	}
}
