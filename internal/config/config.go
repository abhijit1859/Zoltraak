
package config

import (
	
	"time"
)

// Duration is a wrapper around time.Duration to support YAML string parsing
type Duration time.Duration

func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	dur, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*d = Duration(dur)
	return nil
}

func (d Duration) ToDuration() time.Duration {
	return time.Duration(d)
}

// Config represents the top-level configuration tree (Spellbook)
type Config struct {
	Version  string          `yaml:"version"`
	Gateway  GatewayConfig   `yaml:"gateway"`
	Services []ServiceConfig `yaml:"services"`
	Routes   []RouteConfig   `yaml:"routes"`
}

type GatewayConfig struct {
	Name                    string   `yaml:"name"`
	Port                    int      `yaml:"port"`
	AdminPort               int      `yaml:"admin_port"`
	GracefulShutdownTimeout Duration `yaml:"graceful_shutdown_timeout"`
	Environment             string   `yaml:"environment"` // production, development
}

type ServiceConfig struct {
	ID                 string             `yaml:"id"`
	Name               string             `yaml:"name"`
	Discovery          string             `yaml:"discovery"` // static, dynamic
	AureoleServiceName string             `yaml:"aureole_service_name,omitempty"`
	StaticEndpoints    []string           `yaml:"static_endpoints"`
	LoadBalancer       LoadBalancerConfig `yaml:"load_balancer"`
	HealthCheck        HealthCheckConfig  `yaml:"health_check"`
}

type LoadBalancerConfig struct {
	Strategy string `yaml:"strategy"` // round_robin, weighted_round_robin, least_connections
}

type HealthCheckConfig struct {
	Path               string   `yaml:"path"`
	Interval           Duration `yaml:"interval"`
	Timeout            Duration `yaml:"timeout"`
	UnhealthyThreshold int      `yaml:"unhealthy_threshold"`
	HealthyThreshold   int      `yaml:"healthy_threshold"`
}

type RouteConfig struct {
	ID             string              `yaml:"id"`
	Path           string              `yaml:"path"`
	Method         string              `yaml:"method"`
	ServiceID      string              `yaml:"service_id"`
	StripPrefix    string              `yaml:"strip_prefix,omitempty"`
	Timeout        Duration            `yaml:"timeout"`
	RateLimiting   *RateLimitConfig    `yaml:"rate_limiting,omitempty"`
	CircuitBreaker *CircuitBreakerConfig `yaml:"circuit_breaker,omitempty"`
	Retry          *RetryConfig        `yaml:"retry,omitempty"`
}

type RateLimitConfig struct {
	Enabled   bool     `yaml:"enabled"`
	Strategy  string   `yaml:"strategy"` // token_bucket, sliding_window
	Limit     int      `yaml:"limit"`
	Window    Duration `yaml:"window"`
	ClientKey string   `yaml:"client_key"` // ip, api_key
}

type CircuitBreakerConfig struct {
	Enabled          bool     `yaml:"enabled"`
	MinRequests      int32    `yaml:"min_requests"`
	FailureThreshold float64  `yaml:"failure_threshold"`
	CooldownPeriod   Duration `yaml:"cooldown_period"`
	ProbeThreshold   int32    `yaml:"probe_threshold"`
}

type RetryConfig struct {
	Enabled        bool     `yaml:"enabled"`
	MaxAttempts    int      `yaml:"max_attempts"`
	InitialBackoff Duration `yaml:"initial_backoff"`
	MaxBackoff     Duration `yaml:"max_backoff"`
}