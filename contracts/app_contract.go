package contracts

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zerpto/ponodo/cli/contracts"
	"github.com/zerpto/ponodo/config"
	"gorm.io/gorm"
)

type AppContract interface {
	SetupBaseDependencies()
	Run()

	AddCommand(func(app AppContract) contracts.CommandContract)

	SetConfigLoader(*config.Loader)
	GetConfigLoader() *config.Loader
	SetGin(*gin.Engine)
	GetGin() *gin.Engine
	GetDb() *gorm.DB
	SetValidator(*validator.Validate)
	GetValidator() *validator.Validate
}
