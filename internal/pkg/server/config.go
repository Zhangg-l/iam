package server

import (
	"net"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	RecommendedHomeDir   = ".iam"
	RecommendedEnvPrefix = "IAM"
)

type Config struct {
	SecureServing   *SecureServingInfo
	InsecureServing *InsercureServingInfo
	Jwt             *JWTInfo
	Mode            string
	Middlewares     []string
	Healthz         bool
	EnableProfiling bool
	EnableMetrics   bool
}

type Certkey struct {
	CertFile string
	KeyFile  string
}

type SecureServingInfo struct {
	BindAddress string
	BindPort    int
	Certkey     Certkey
}

func (s *SecureServingInfo) Address() string {
	return net.JoinHostPort(s.BindAddress, strconv.Itoa(s.BindPort))
}

type JWTInfo struct {
	// default to "iam jwt"
	Realm string
	// default to empty
	Key string
	// default one hour
	Timeout time.Duration
	// default to zero
	MaxRefresh time.Duration
}

type InsercureServingInfo struct {
	Address string
}

func NewConfig() *Config {
	return &Config{
		Healthz:         true,
		Mode:            gin.ReleaseMode,
		Middlewares:     []string{},
		EnableMetrics:   true,
		EnableProfiling: true,
		Jwt: &JWTInfo{
			Realm:      "iam Jwt",
			Timeout:    1 * time.Hour,
			MaxRefresh: 1 * time.Hour,
		},
	}
}

type CompletedConfig struct {
	*Config
}

func (c *CompletedConfig) New() error {

	return nil
}

func LoadConfig(cfg string, defaultName string) {

}
