package ernos

import (
	"net/http"
	"strings"
)

type Ernos struct {
	Status    int    `json:"status,omitempty"`
	Message   string `json:"message,omitempty"`
	Code      string `json:"code,omitempty"`
	Attribute any    `json:"attribute,omitempty"`
}

func (e Ernos) Error() string {
	return e.Message
}

func HasAlreadyExists(value ...string) error {
	msg := M.DUPLICATE_RECORD
	if len(value) > 0 {
		msg = strings.Join(value, "") + " has already exists"
	}

	return Ernos{
		Status:  http.StatusConflict,
		Message: msg,
		Code:    C.DUPLICATE_RECORD,
	}
}

func NotFound(value ...string) error {
	msg := M.RECORD_NOTFOUND
	if len(value) > 0 {
		msg = strings.Join(value, "") + " not found"
	}

	return Ernos{
		Message: msg,
		Code:    C.RECORD_NOTFOUND,
	}
}

func Forbidden(value string) error {
	msg := http.StatusText(http.StatusForbidden)
	if value != "" {
		msg = value
	}

	return Ernos{
		Message: msg,
		Code:    C.RECORD_NOTFOUND,
	}
}

func Unauthorized(value ...string) error {
	msg := M.UNAUTHORIZED
	if len(value) > 0 {
		msg = strings.Join(value, "")
	}

	return Ernos{
		Status:  http.StatusUnauthorized,
		Message: msg,
		Code:    C.UNAUTHORIZED,
	}
}

func InternalServerError(value ...string) error {
	msg := M.INTERNAL_SERVER_ERROR
	if len(value) > 0 {
		msg = strings.Join(value, "")
	}

	return Ernos{
		Status:  http.StatusInternalServerError,
		Message: msg,
		Code:    C.INTERNAL_SERVER_ERROR,
	}
}

func Other(e Ernos) error {
	return e
}
