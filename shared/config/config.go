package config

import (
	"fmt"
	"os"
	"sync"

	SharedError "github.com/github-profile/go-nats/shared/error"
	"github.com/spf13/viper"
)

type (

	// ImmutableConfigInterface is an interface represent methods in config
	ImmutableConfigInterface interface {
		GetPort() int
		GetNATSHost() string
	}

	// im is a struct to mapping the structure of related value model
	im struct {
		Port     int    `mapstructure:"PORT"`
		NATSHost string `mapstructure:"NATS_HOST"`
	}
)

func (i *im) GetPort() int {
	return i.Port
}

func (i *im) GetNATSHost() string {
	return i.NATSHost
}

var (
	imOnce    sync.Once
	myEnv     map[string]string
	immutable im
)

// NewImmutableConfig is a factory that return of its config implementation
func NewImmutableConfig() ImmutableConfigInterface {
	imOnce.Do(func() {
		v := viper.New()
		appEnv, exists := os.LookupEnv("APP_ENV")

		if exists {
			if appEnv == "staging" {
				v.SetConfigName("app.config.staging")
			} else if appEnv == "production" {
				v.SetConfigName("app.config.prod")
			} else if appEnv == "development" {
				v.SetConfigName("app.config.dev")
			}
		} else {
			appEnv = "development"
			v.SetConfigName("app.config.dev")
		}
		fmt.Printf("Reading app_env: %s\n", appEnv)

		v.AddConfigPath(".")
		v.SetEnvPrefix("PROMOTION")
		v.AutomaticEnv()

		if err := v.ReadInConfig(); err != nil {
			SharedError.Wrap(500, "[PROMOTION-SYS-001]", err, "[CONFIG][missing] Failed to read app.config.* file", "failed")
		}

		v.Unmarshal(&immutable)
	})

	return &immutable
}
