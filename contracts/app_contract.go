package contracts

import (
	"github.com/gin-gonic/gin"
	"github.com/zerpto/ponodo/config"
	"gorm.io/gorm"
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
