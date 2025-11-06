package contracts

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	clicontracts "github.com/zerpto/ponodo/cli/contracts"
	"github.com/zerpto/ponodo/config"
	"gorm.io/gorm"
)

// AppContract defines the interface for the main application structure.
// It provides methods to manage application lifecycle, dependencies, and
// core services like configuration, database, HTTP server, and validation.
//
//go:generate mockgen -source=$GOFILE -destination=./mocks/mock_app_contract.go -package=mocks
type AppContract interface {
	SetupBaseDependencies()
	Run()

	AddCommand(func(app AppContract) clicontracts.CommandContract)

	SetConfigLoader(*config.Loader)
	GetConfigLoader() *config.Loader
	SetGin(*gin.Engine)
	GetGin() *gin.Engine
	GetDb() *gorm.DB
	SetValidator(*validator.Validate)
	GetValidator() *validator.Validate
}
