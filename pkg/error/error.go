package custom_error

import "github.com/ansxy/golang-boilerplate-gin/pkg/constant"

type ErrorContext struct {
	Code     int
	Message  string
	Function string
}

type CustomeError struct {
	ErrorContext *ErrorContext
}

func (e *CustomeError) Error() string {
	if e.ErrorContext.Message == "" {
		e.ErrorContext.Message = constant.ErrorMessageMap[e.ErrorContext.Code]
	}

	return constant.ErrorMessageMap[constant.DefaultUnhandledErrorCode]
}

func SetCustomeError(cte *ErrorContext) *CustomeError {
	return &CustomeError{ErrorContext: cte}
}
