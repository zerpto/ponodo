package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Loader struct {
	Config *Config
}

func (c *Loader) loadFromEnvironmentVariable() error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.AutomaticEnv()
	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return nil
}

func (c *Loader) BindTo(config *Config) error {
	c.Config = config

	if err := viper.Unmarshal(c.Config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %v", err)
	}

	if err := viper.Unmarshal(&c.Config.DB); err != nil {
		return fmt.Errorf("failed to unmarshal DB: %v", err)
	}
	return nil
}

func NewLoader() (*Loader, error) {
	loader := Loader{}
	err := loader.loadFromEnvironmentVariable()
	if err != nil {
		return nil, err
	}
	return &loader, nil
}
