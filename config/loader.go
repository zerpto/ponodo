package config

import (
	"strings"

	"github.com/zerpto/ponodo/config/contracts"

	"github.com/spf13/viper"
)

// Loader handles loading and managing application configuration.
// It reads configuration from environment variables and provides access
// to the loaded configuration through the ConfigContract interface.
type Loader struct {
	Config contracts.ConfigContract
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

//func (c *Loader) BindTo(config contracts.ConfigContract) error {
//	c.Config = config
//
//	if err := viper.Unmarshal(c.Config); err != nil {
//		return fmt.Errorf("failed to unmarshal config: %v", err)
//	}
//
//	if err := viper.Unmarshal(&c.Config.DB); err != nil {
//		return fmt.Errorf("failed to unmarshal DB: %v", err)
//	}
//	return nil
//}

// NewLoader creates a new configuration loader instance.
// It initializes the loader and loads configuration from environment variables.
// Returns an error if the configuration cannot be loaded.
func NewLoader() (*Loader, error) {
	loader := Loader{}
	err := loader.loadFromEnvironmentVariable()
	if err != nil {
		return nil, err
	}
	return &loader, nil
}
