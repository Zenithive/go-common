package errorx

import (
	"errors"
)

type Error struct {
	Code   string
	Msg    string
	Status int
	Err    error
	Detail any
}

func (e *Error) Error() string {
	return e.Msg
}

func (e *Error) Unwrap() error {
	return e.Err
}

func From(err error) *Error {
	if err == nil {
		return nil
	}

	var appErr *Error

	if errors.As(err, &appErr) {
		return appErr
	}

	return Internal(err)
}
