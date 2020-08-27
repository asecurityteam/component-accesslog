package accesslog

import (
	"context"
	"net/http"

	"github.com/asecurityteam/transport"
)

// Config modifies the behavior of the access logs.
type Config struct{}

// Name of the config root.
func (*Config) Name() string {
	return "accesslog"
}

// Component is a logging plugin.
type Component struct{}

// NewComponent initializes a Component with default values.
func NewComponent() *Component {
	return &Component{}
}

// Settings generates a config populated with defaults.
func (m *Component) Settings() *Config {
	return &Config{}
}

// New generates the middleware.
func (*Component) New(ctx context.Context, conf *Config) (func(http.RoundTripper) http.RoundTripper, error) {
	return func(next http.RoundTripper) http.RoundTripper {
		return transport.NewAccessLog()(next)
	}, nil
}
