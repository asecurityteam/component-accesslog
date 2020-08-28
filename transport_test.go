package accesslog

import (
	"context"
	"net/http"
	"testing"

	"github.com/asecurityteam/settings"
	"github.com/stretchr/testify/assert"
)

func TestConfig_Name(t *testing.T) {
	Config := Config{}
	assert.Equal(t, "accesslog", Config.Name())
}

func TestAccessLog_Settings(t *testing.T) {
	accessLogComponent := &Component{}
	assert.IsType(t, &Config{}, accessLogComponent.Settings())
}

func TestAccessLog_AccessLog(t *testing.T) {
	component := NewComponent()
	assert.IsType(t, &Component{}, component)
}

func TestAccessLog_New(t *testing.T) {
	accessLogConfig := Config{}
	accessLogComponent := Component{}
	wrapper, err := accessLogComponent.New(context.Background(), &accessLogConfig)
	assert.IsType(t, func(next http.RoundTripper) http.RoundTripper { return nil }, wrapper)
	assert.Nil(t, err)
	transport := wrapper(http.DefaultTransport)
	assert.NotNil(t, transport)
}

func TestNew(t *testing.T) {
	src := settings.NewMapSource(map[string]interface{}{
		"accesslog": map[string]interface{}{},
	})
	wrapper, err := New(context.Background(), src)
	assert.Nil(t, err)
	assert.IsType(t, wrapper, func(next http.RoundTripper) http.RoundTripper { return nil }, wrapper)
}
