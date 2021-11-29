package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Build information -ldflags .
const (
	version    string = "dev"
	commitHash string = "-"
)

var cfg *Config

// GetConfigInstance returns service config
func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

// Project - contains all parameters project information.
type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	ServiceName string `yaml:"serviceName"`
	Version     string
	CommitHash  string
}

// Metrics - contains all parameters metrics information.
type Metrics struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
	Path string `yaml:"path"`
}

// Telemetry log path
type Telemetry struct {
	GraylogPath string `yaml:"graylogPath"`
}

// Jaeger - contains all parameters metrics information.
type Jaeger struct {
	Service      string  `yaml:"service"`
	Host         string  `yaml:"host"`
	Port         string  `yaml:"port"`
	SamplerType  string  `yaml:"samplerType"`
	SamplerParam float64 `yaml:"samplerParam"`
}

// Status config for service.
type Status struct {
	Port          int    `yaml:"port"`
	Host          string `yaml:"host"`
	VersionPath   string `yaml:"versionPath"`
	LivenessPath  string `yaml:"livenessPath"`
	ReadinessPath string `yaml:"readinessPath"`
}

type GrpcLgcGroup struct {
	Port    int    `yaml:"port"`
	Timeout int64  `yaml:"timeout"`
	Host    string `yaml:"host"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project             Project      `yaml:"project"`
	GrpcLgcGroupApi     GrpcLgcGroup `yaml:"grpc-lgc-group-api"`
	GrpcLgcGroupFacade  GrpcLgcGroup `yaml:"grpc-lgc-group-facade"`
	MetricsRetranslator Metrics      `yaml:"metricsRetranslator"`
	MetricsGrpc         Metrics      `yaml:"metrics"`
	Telemetry           Telemetry    `yaml:"telemetry"`
	Jaeger              Jaeger       `yaml:"jaeger"`
	Status              Status       `yaml:"status"`
}

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(filePath string) (err error) {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Unable to close file: %s", err)
		}
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	cfg.Project.Version = version
	cfg.Project.CommitHash = commitHash

	return nil
}
