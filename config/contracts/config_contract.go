package contracts

// ConfigContract defines the interface for application configuration access.
// Implementations of this interface provide methods to retrieve application
// settings including app name, environment, debug mode, and database configuration.
//
//go:generate mockgen -source=$GOFILE -destination=./mocks/mock_config_contract.go -package=mocks
type ConfigContract interface {
	GetApp() string
	GetEnv() string
	GetDebug() bool

	GetDb() DbConfigContract
}
