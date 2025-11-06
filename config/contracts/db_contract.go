package contracts

// DbConfigContract defines the interface for database configuration access.
// Implementations of this interface provide methods to retrieve database
// connection parameters such as host, port, user, password, and database name.
//
//go:generate mockgen -source=$GOFILE -destination=./mocks/mock_db_contract.go -package=mocks
type DbConfigContract interface {
	GetHost() string
	GetPort() string
	GetUser() string
	GetPassword() string
	GetDatabase() string
}
