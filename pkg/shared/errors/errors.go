package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
	Type   string `json:"type"`
	Status int    `json:"-"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s: %s", e.Code, e.Type, e.Detail)
}

func New(code int, t string, detail string, status ...int) *Error {
	e := &Error{
		Code:   code,
		Type:   t,
		Detail: detail,
		Status: http.StatusInternalServerError,
	}

	if len(status) > 0 {
		e.Status = status[0]
	}

	return e
}

func NewDynamic(code int, t string) func(detail string, status ...int) *Error {
	return func(detail string, status ...int) *Error {
		e := &Error{
			Code:   code,
			Type:   t,
			Detail: detail,
			Status: http.StatusInternalServerError,
		}

		if len(status) > 0 {
			e.Status = status[0]
		}

		return e
	}
}
