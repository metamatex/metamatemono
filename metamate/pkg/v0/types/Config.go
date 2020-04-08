package types

import (
	"github.com/metamatex/metamate/gen/v0/sdk"
)

type Config struct {
	DiscoverySvc         sdk.Service       `yaml:"discoverySvc,omitempty"`
	Endpoints            EndpointsConfig   `yaml:"endpoints,omitempty"`
	Host                 HostConfig        `yaml:"host,omitempty"`
	Log                  LogConfig         `yaml:"log,omitempty"`
	Virtual              VirtualConfig     `yaml:"virtual,omitempty"`
}

type VirtualConfig struct {
	Services []VirtualSvc
}

type LogConfig struct {
	Http bool
	Internal InternalLogConfig
}

// stage, type, format
type InternalLogConfig map[string]map[string]string

type EndpointsConfig struct {
	Config           ConfigEndpointConfig           `yaml:"config,omitempty"`
	Prometheus       PrometheusEndpointConfig       `yaml:"prometheus,omitempty"`
	Debug            DebugEndpointConfig            `yaml:"debug,omitempty"`
	Graphql          GraphqlEndpointConfig          `yaml:"graphql,omitempty"`
	GraphiqlExplorer GraphiqlExplorerEndpointConfig `yaml:"graphiqlExplorer,omitempty"`
	HttpJson         HttpJsonEndpoint               `yaml:"httpJson,omitempty"`
}

type ConfigEndpointConfig struct {
	On bool `yaml:"on,omitempty"`
}

type PrometheusEndpointConfig struct {
	On   bool   `yaml:"on,omitempty"`
}

type DebugEndpointConfig struct {
	On bool `yaml:"on,omitempty"`
}

type HostConfig struct {
	AllowedOrigins []string `yaml:"allowedOrigins,omitempty"`
}

type GraphiqlExplorerEndpointConfig struct {
	On           bool   `yaml:"on,omitempty"`
	DefaultQuery string `yaml:"defaultQuery,omitempty"`
}

type GraphqlEndpointConfig struct {
	On             bool     `yaml:"on,omitempty"`
	PlaygroundPath string   `yaml:"playgroundPath,omitempty"`
}

type HttpJsonEndpoint struct {
	On   bool   `yaml:"on,omitempty"`
}
