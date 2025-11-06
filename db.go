package ponodo

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewGormConnection creates a new GORM database connection to PostgreSQL.
// It constructs the connection string from the provided parameters and
// returns a configured GORM DB instance ready for use.
func NewGormConnection(host, port, user, password, dbName string) *gorm.DB {
	if port == "" {
		port = "5432"
	}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
