package contracts

import (
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/zerpto/ponodo/config"
)

type AppContract interface {
	SetupBaseDependencies()
	Run()

	GetConfigLoader() *config.Loader
	AddCommand(func(app AppContract) CommandContract)

	SetGin(*gin.Engine)
	GetGin() *gin.Engine
	GetDb() *gorm.DB
	SetValidator(*validator.Validate)
	GetValidator() *validator.Validate
}
