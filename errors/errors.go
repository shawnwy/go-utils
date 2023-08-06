package errors

import (
	"errors"
	"fmt"
)

type ierror interface {
	Code() int
}

// New returns an error that formats as the given text. Each call to New returns a distinct error value even if the text is identical.
func New(message string) error { return errors.New(message) }

//Is reports whether any error in chain of {err} matches target.
// The chain consists of err itself followed by the sequence of errors obtained by repeatedly calling Unwrap.
// An error is considered to match a target if it is equal to that target or if it implements a method Is(error) bool such that Is(target) returns true.
// An error type might provide an Is method, so it can be treated as equivalent to an existing error. For example, if MyError defines
func Is(err, target error) bool { return errors.Is(err, target) }

// As finds the first error in chain of {err} that matches target, and if so, sets target to that error value and returns true. Otherwise, it returns false.
// The chain consists of err itself followed by the sequence of errors obtained by repeatedly calling Unwrap.
// An error matches target if the error's concrete value is assignable to the value pointed to by target, or if the error has a method As(interface{}) bool such that As(target) returns true. In the latter
func As(err error, target interface{}) bool { return errors.As(err, target) }

// withCode is an error that has a message and a specific code
type withCode struct {
	error
	msg  string
	code int
}

func (c *withCode) Error() string { return fmt.Sprintf("%d - %s", c.code, c.msg) }
func (c *withCode) Unwrap() error { return c.error }
func (c *withCode) Code() int     { return c.code }

// NewWithCode returns an error with supplied message and a specific code.
func NewWithCode(code int, message string) error {
	return &withCode{
		msg:  message,
		code: code,
	}
}

// GetCode will try to get 'code' of error
// if the error is not ierror, it will return 0
// otherwise, it will return the 'code'
func GetCode(err error) int {
	u, ok := err.(interface {
		Code() int
	})
	if !ok {
		return 0
	}
	return u.Code()
}

// Wrap returns an error with additional message
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return &withCode{
		err,
		fmt.Sprintf("'%s' caused by: %s)", message, err.Error()),
		GetCode(err),
	}
}

// Wrapf returns an error with additional message
func Wrapf(err error, message string, args ...interface{}) error {
	return Wrap(err, fmt.Sprintf(message, args...))
}

// With will combine new error with the cause
// and the combined error will have the code of cause error
func With(cause, err error) error {
	if cause == nil {
		return err
	}
	return &withCode{
		err,
		fmt.Sprintf("'%s' caused by: %s)", err.Error(), cause.Error()),
		GetCode(cause),
	}
}
