package herror

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	herr := NewHTTPErrorImpl(1, "Any error", "Reason of any error", nil)
	assert.Equal(t, 1, herr.Code())
	assert.Equal(t, "Any error", herr.PublicMessage())
	assert.Equal(t, "Reason of any error", herr.PrivateMessage())
	assert.False(t, herr.Is4XX())
	assert.Regexp(t, `^.+herror_test\.go:10$`, herr.Call())
	assert.Nil(t, herr.Error())
}

func TestNotFound(t *testing.T) {
	herr := NewNotFound("Any error", "Reason of any error", nil)
	assert.Equal(t, 404, herr.Code())
	assert.Equal(t, "NotFound: Any error", herr.PublicMessage())
	assert.Equal(t, "Reason of any error", herr.PrivateMessage())
	assert.True(t, herr.Is4XX())
	assert.Regexp(t, `^.+herror_test\.go:20$`, herr.Call())
	assert.Nil(t, herr.Error())
}

func TestInternalServerError(t *testing.T) {
	herr := NewInternalServerError("Any error", "Reason of any error", nil)
	assert.Equal(t, 500, herr.Code())
	assert.Equal(t, "InternalServerError: Any error", herr.PublicMessage())
	assert.Equal(t, "Reason of any error", herr.PrivateMessage())
	assert.False(t, herr.Is4XX())
	assert.Regexp(t, `^.+herror_test\.go:30$`, herr.Call())
	assert.Nil(t, herr.Error())
}

func TestUnauthorizedBadAccessToken(t *testing.T) {
	herr := NewUnauthorizedBadAccessToken(nil)
	assert.Equal(t, 401, herr.Code())
	assert.Equal(t, "Bad access token", herr.PublicMessage())
	assert.Equal(t, "Bad access token", herr.PrivateMessage())
	assert.True(t, herr.Is4XX())
	assert.Regexp(t, `^.+herror_test\.go:40$`, herr.Call())
	assert.Nil(t, herr.Error())
}

func TestBindJSONError(t *testing.T) {
	herr := NewBindJSONError(nil)
	assert.Equal(t, 400, herr.Code())
	assert.Equal(t, "Cannot bind body", herr.PublicMessage())
	assert.Equal(t, "Cannot bind body", herr.PrivateMessage())
	assert.True(t, herr.Is4XX())
	assert.Regexp(t, `^.+herror_test\.go:50$`, herr.Call())
	assert.Nil(t, herr.Error())
}

func TestInvalidParameterError(t *testing.T) {
	herr := NewInvalidParameterError("pri", "pub", nil)
	assert.Equal(t, 400, herr.Code())
	assert.Equal(t, "pri", herr.PublicMessage())
	assert.Equal(t, "pub", herr.PrivateMessage())
	assert.True(t, herr.Is4XX())
	assert.Regexp(t, `^.+herror_test\.go:60$`, herr.Call())
	assert.Nil(t, herr.Error())
}

func TestConflictError(t *testing.T) {
	herr := NewConflictError("pri", "pub", nil)
	assert.Equal(t, 409, herr.Code())
	assert.Equal(t, "pri", herr.PublicMessage())
	assert.Equal(t, "pub", herr.PrivateMessage())
	assert.True(t, herr.Is4XX())
	assert.Regexp(t, `^.+herror_test\.go:70$`, herr.Call())
	assert.Nil(t, herr.Error())
}
