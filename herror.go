package herror

import (
	"fmt"
	"net/http"
	"runtime"
)

func getCallInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return fmt.Sprintf("unknown:unknown")
}

const (
	// StatusAgentExpired ...
	StatusAgentExpired = 450
	// StatusLoginExpired ...
	StatusLoginExpired = 451
)

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
		call:           getCallInfo(2),
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
		call:           getCallInfo(2),
		err:            err,
	}
}

// NewInternalServerError ...
func NewInternalServerError(publicMessage, privateMessage string, err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusInternalServerError,
		publicMessage:  fmt.Sprintf("InternalServerError: %s", publicMessage),
		privateMessage: privateMessage,
		call:           getCallInfo(2),
		err:            err,
	}
}

// NewUnauthorizedBadAccessToken ...
func NewUnauthorizedBadAccessToken(err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusUnauthorized,
		publicMessage:  "Bad access token",
		privateMessage: "Bad access token",
		call:           getCallInfo(2),
		err:            err,
	}
}

// NewBindJSONError ...
func NewBindJSONError(err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusBadRequest,
		publicMessage:  "Cannot bind body",
		privateMessage: "Cannot bind body",
		call:           getCallInfo(2),
		err:            err,
	}
}

// NewInvalidParameterError ...
func NewInvalidParameterError(pri, pub string, err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusBadRequest,
		publicMessage:  pri,
		privateMessage: pub,
		call:           getCallInfo(2),
		err:            err,
	}
}

// NewConflictError ...
func NewConflictError(pri, pub string, err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusConflict,
		publicMessage:  pri,
		privateMessage: pub,
		call:           getCallInfo(2),
		err:            err,
	}
}

// NewLoginError ...
func NewLoginError(pri string, err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusForbidden,
		publicMessage:  "Login process is failed",
		privateMessage: pri,
		call:           getCallInfo(2),
		err:            err,
	}
}

// NewAgentExpiredError ...
func NewAgentExpiredError(pri string, err error) HTTPError {
	return HTTPErrorImpl{
		code:           StatusAgentExpired,
		publicMessage:  "Agent token is expired",
		privateMessage: pri,
		call:           getCallInfo(2),
		err:            err,
	}
}

// NewLoginExpiredError ...
func NewLoginExpiredError(pri string, err error) HTTPError {
	return HTTPErrorImpl{
		code:           StatusLoginExpired,
		publicMessage:  "Login token is expired",
		privateMessage: pri,
		call:           getCallInfo(2),
		err:            err,
	}
}

// NewUnauthorizedError ...
func NewUnauthorizedError(pub, pri string, err error) HTTPError {
	return HTTPErrorImpl{
		code:           http.StatusUnauthorized,
		publicMessage:  pub,
		privateMessage: pri,
		call:           getCallInfo(2),
		err:            err,
	}
}
