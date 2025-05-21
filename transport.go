package accesslog

import (
	"context"
	"net/http"

	"github.com/asecurityteam/settings/v2"
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

// New is the top-level entrypoint for creating an `http.Transport` decorator
// that emits an access logline on every `RoundTrip`.
//
// Useful when configuring a `Component` outside of the hierarchy of a
// surrounding appliation.
func New(ctx context.Context, source settings.Source) (func(http.RoundTripper) http.RoundTripper, error) {
	var dst func(http.RoundTripper) http.RoundTripper
	_ = settings.NewComponent(ctx, source, NewComponent(), &dst)
	return dst, nil
}
