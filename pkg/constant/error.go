package constant

import "net/http"

const (
	DefaultUnhandledErrorCode = 1000 + iota
	DefaultNotFoundErrorCode
	DefaultInvalidRequestErrorCode
	DefaultDuplicateErrorCode
	DefaultUnauthorizedErrorCode
	DefaultValidateErrorCode
)

var ErrorMessageMap = map[int]string{
	DefaultUnhandledErrorCode:      "Unhandled error",
	DefaultNotFoundErrorCode:       "Not found",
	DefaultInvalidRequestErrorCode: "Invalid request",
	DefaultDuplicateErrorCode:      "Duplicate",
	DefaultUnauthorizedErrorCode:   "Unauthorized",
	DefaultValidateErrorCode:       "Validate error",
}

var ErrorCodeMap = map[int]int{
	DefaultDuplicateErrorCode:      http.StatusConflict,
	DefaultUnauthorizedErrorCode:   http.StatusUnauthorized,
	DefaultValidateErrorCode:       http.StatusBadRequest,
	DefaultNotFoundErrorCode:       http.StatusNotFound,
	DefaultInvalidRequestErrorCode: http.StatusUnprocessableEntity,
}
