package contracts

// DbConfigContract defines the interface for database configuration access.
// Implementations of this interface provide methods to retrieve database
// connection parameters such as host, port, user, password, and database name.
type DbConfigContract interface {
	GetHost() string
	GetPort() string
	GetUser() string
	GetPassword() string
	GetDatabase() string
}

// ConfigContract defines the interface for application configuration access.
// Implementations of this interface provide methods to retrieve application
// settings including app name, environment, debug mode, and database configuration.
type ConfigContract interface {
	GetApp() string
	GetEnv() string
	GetDebug() bool

	GetDb() DbConfigContract
}
