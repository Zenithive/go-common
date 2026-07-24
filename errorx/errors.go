package errorx

import "net/http"

func BadReq(msg string) *Error {
	return &Error{
		Code:   CodeBadReq,
		Msg:    msg,
		Status: http.StatusBadRequest,
	}
}

func Validation(detail any) *Error {
	return &Error{
		Code:   CodeValid,
		Msg:    "Request validation failed",
		Status: http.StatusUnprocessableEntity,
		Detail: detail,
	}
}

func Unauth(msg string) *Error {
	return &Error{
		Code:   CodeUnauth,
		Msg:    msg,
		Status: http.StatusUnauthorized,
	}
}

func Forbidden(msg string) *Error {
	return &Error{
		Code:   CodeForbidden,
		Msg:    msg,
		Status: http.StatusForbidden,
	}
}

func NotFound(msg string) *Error {
	return &Error{
		Code:   CodeNotFound,
		Msg:    msg,
		Status: http.StatusNotFound,
	}
}

func Conflict(msg string) *Error {
	return &Error{
		Code:   CodeConflict,
		Msg:    msg,
		Status: http.StatusConflict,
	}
}

func Internal(err error) *Error {
	return &Error{
		Code:   CodeInternal,
		Msg:    "Internal server error",
		Status: http.StatusInternalServerError,
		Err:    err,
	}
}
