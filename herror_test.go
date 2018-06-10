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