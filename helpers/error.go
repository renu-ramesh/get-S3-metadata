package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ErrNotImplemented is returned when the requested method has not been implemented
var ErrNotImplemented = errors.New("not implemented")

// ErrNotFound is returned when the requested resource was not found
var ErrNotFound = errors.New("requested resource not found")

// ErrBadRequest is returned when the request could not be completed due to an error in the request itself
var ErrBadRequest = errors.New("unable to complete the request")

// ErrUnknown is returned when the server encountered an error preventing the request from being completed
var ErrUnknown = errors.New("unknown server error")

// ErrUnauthorized is returned when the request could not be completed because the principal was not authorized
var ErrUnauthorized = errors.New("unauthorized")

//Error validation Error
var ErrValidationError = errors.New("validation error")

type Error interface {
	error
	ToJson() string
	Err() error
	Msg() string
}

var _ error = &serviceError{}

type serviceError struct {
	err  error
	msg  string
	code string
}

func (e serviceError) Error() string {
	return fmt.Sprintf("%s; %s", e.err, e.msg)
}

func (e serviceError) Err() error {
	return e.err
}

func (e serviceError) Msg() string {
	return e.msg
}

func (e serviceError) ToJson() string {
	type S struct {
		Err  string `json:"error"`
		Msg  string `json:"msg"`
		Code string `json:"code"`
	}

	b, _ := json.Marshal(S{e.err.Error(), e.msg, e.code})
	return string(b)
}

func NewError(err error, msg string, args ...interface{}) Error {
	return serviceError{
		err: err,
		msg: msg,
	}
}
