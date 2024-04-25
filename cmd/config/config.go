package config

import (
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

const (
	configName   = "monitor-creator"
	configPrefix = "monitor-creator"
)

var (
	instance             *ConfigManager
	once                 sync.Once
	monitorCreatorConfig MonitorCreatorConfig
)

type ConfigManager struct {
	viper *viper.Viper
}

func Initialize() *ConfigManager {
	once.Do(func() {
		instance = &ConfigManager{
			viper: viper.New(),
		}
	})
	return instance
}

func Get() MonitorCreatorConfig {
	return monitorCreatorConfig
}

func LoadMonitorCreatorConfig() {
	Initialize()

	viper.SetEnvPrefix(configPrefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName(configName)
	viper.AddConfigPath("../config")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	monitorCreatorConfig = MonitorCreatorConfig{
		LogLevel: viper.GetString("log_level"),
		Auth: AuthConfig{
			Kind:       viper.GetString("auth.kind"),
			Middleware: viper.GetString("auth.middleware"),
		},
		DatadogURL: viper.GetString("datadog_api_url"),
		ClientAuth: HTTPClientCredential{
			Token:    os.Getenv("FURY_TOKEN"),
			ClientID: os.Getenv("SECRET_CLIENT_ID"),
			SecretID: os.Getenv("SECRET_SECRET_ID"),
		},
	}
}

type MonitorCreatorConfig struct {
	LogLevel   string
	Auth       AuthConfig
	ClientAuth HTTPClientCredential
	DatadogURL string
}

type AuthConfig struct {
	Kind       string
	Middleware string
}

type HTTPClientCredential struct {
	Token    string
	ClientID string
	SecretID string
}
