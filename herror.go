package herror

import (
	"fmt"
	"net/http"
	"runtime"
)

func getCallInfo() string {
	skip := 2
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return fmt.Sprintf("unknown:unknown")
}

// HTTPError ...
type HTTPError interface {
	Code() int
	PrivateMessage() string
	PublicMessage() string
	Is4XX() bool
	Error() error
	Call() string
	String() string
}

// HTTPErrorImpl ...
type HTTPErrorImpl struct {
	code           int
	privateMessage string
	publicMessage  string
	err            error
	call           string
}

// NewHTTPErrorImpl ...
func NewHTTPErrorImpl(code int, publicMessage, privateMessage string, err error) *HTTPErrorImpl {
	return &HTTPErrorImpl{
		code:           code,
		privateMessage: privateMessage,
		publicMessage:  publicMessage,
		err:            err,
		call:           getCallInfo(),
	}
}

// Is4XX ...
func (err HTTPErrorImpl) Is4XX() bool {
	return 400 <= err.code && err.code < 500
}

// Code ...
func (err HTTPErrorImpl) Code() int {
	return err.code
}

// PrivateMessage ...
func (err HTTPErrorImpl) PrivateMessage() string {
	return err.privateMessage
}

// PublicMessage ...
func (err HTTPErrorImpl) PublicMessage() string {
	return err.publicMessage
}

// Error ...
func (err HTTPErrorImpl) Error() error {
	return err.err
}

// Call ...
func (err HTTPErrorImpl) Call() string {
	return err.call
}

// String ...
func (err HTTPErrorImpl) String() string {
	return fmt.Sprintf(
		"Private[%s] Public[%s] Code[%d] Call[%s] Err[%+v]",
		err.privateMessage, err.publicMessage, err.code, err.call, err.err,
	)
}

// NewNotFound ...
func NewNotFound(publicMessage, privateMessage string, err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusNotFound,
		publicMessage:  fmt.Sprintf("NotFound: %s", publicMessage),
		privateMessage: privateMessage,
		call:           getCallInfo(),
		err:            err,
	}
}

// NewInternalServerError ...
func NewInternalServerError(publicMessage, privateMessage string, err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusInternalServerError,
		publicMessage:  fmt.Sprintf("InternalServerError: %s", publicMessage),
		privateMessage: privateMessage,
		call:           getCallInfo(),
		err:            err,
	}
}

// NewUnauthorizedBadAccessToken ...
func NewUnauthorizedBadAccessToken(err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusUnauthorized,
		publicMessage:  "Bad access token",
		privateMessage: "Bad access token",
		call:           getCallInfo(),
		err:            err,
	}
}

// NewBindJSONError ...
func NewBindJSONError(err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusBadRequest,
		publicMessage:  "Cannot bind body",
		privateMessage: "Cannot bind body",
		call:           getCallInfo(),
		err:            err,
	}
}

// NewInvalidParameterError ...
func NewInvalidParameterError(pri, pub string, err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusBadRequest,
		publicMessage:  pri,
		privateMessage: pub,
		call:           getCallInfo(),
		err:            err,
	}
}

// NewConflictError ...
func NewConflictError(pri, pub string, err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusConflict,
		publicMessage:  pri,
		privateMessage: pub,
		call:           getCallInfo(),
		err:            err,
	}
}

// NewLoginError ...
func NewLoginError(pri string, err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusForbidden,
		publicMessage:  "Login process is failed",
		privateMessage: pri,
		call:           getCallInfo(),
		err:            err,
	}
}
