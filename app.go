package ponodo

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	clicontracts "github.com/zerpto/ponodo/cli/contracts"
	"github.com/zerpto/ponodo/contracts"

	//"github.com/zerpto/template-backend-go/src/routers"
	"gorm.io/gorm"

	"github.com/zerpto/ponodo/cli"
	"github.com/zerpto/ponodo/config"
)

type App struct {
	ConfigLoader *config.Loader
	DB           *gorm.DB
	Command      *cobra.Command
	Gin          *gin.Engine
	Validator    *validator.Validate
}

func (app *App) SetConfigLoader(loader *config.Loader) {
	app.ConfigLoader = loader
}

func (app *App) SetValidator(validate *validator.Validate) {
	app.Validator = validate
}

func (app *App) GetValidator() *validator.Validate {
	return app.Validator
}

func (app *App) GetDb() *gorm.DB {
	return app.DB
}

func (app *App) GetGin() *gin.Engine {
	return app.Gin
}

func (app *App) SetGin(engine *gin.Engine) {
	app.Gin = engine
}

func (app *App) SetupBaseDependencies() {
	//app.setupConfig()
	app.setupLogger()
	//app.setupDatabaseConnection()
	app.setupModel()
}

func (app *App) AddCommand(f func(app contracts.AppContract) clicontracts.CommandContract) {
	rootCmd := app.Command

	command := f(app)

	rootCmd.AddCommand(&cobra.Command{
		Use:     command.Use(),
		Short:   command.Short(),
		Long:    command.Long(),
		Example: command.Example(),
		Run: func(cobra *cobra.Command, args []string) {
			command.Run(cobra, args)
		},
	})
}

func (app *App) Run() {
	cliApp := cli.NewCli(app)
	//routers.NewCliRouter(cliApp)
	cliApp.Run()
}

func (app *App) GetConfigLoader() *config.Loader {
	return app.ConfigLoader
}

//func (app *App) setupConfig() {
//	cfg, err := config.NewLoader()
//	if err != nil {
//		panic(err)
//	}
//	err = cfg.BindTo(&config.Config{}) // todo: move this outside
//	if err != nil {
//		panic(err)
//	}
//	app.ConfigLoader = cfg
//}

func (app *App) setupDatabaseConnection() {

}

func (app *App) setupLogger() {
	NewLogger()
}

func (app *App) setupModel() {
	cfg := app.ConfigLoader.Config
	dbCfg := cfg.GetDb()
	host := dbCfg.GetHost()
	port := dbCfg.GetPort()
	user := dbCfg.GetUser()
	password := dbCfg.GetPassword()
	dbName := dbCfg.GetDatabase()

	db := NewGormConnection(host, port, user, password, dbName)
	app.DB = db
}

func NewApp() contracts.AppContract {
	return &App{}
}
