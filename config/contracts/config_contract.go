package contracts

type DbConfigContract interface {
	GetHost() string
	GetPort() string
	GetUser() string
	GetPassword() string
	GetDatabase() string
}

type ConfigContract interface {
	GetApp() string

	GetDb() DbConfigContract
}
