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

// App represents the main application structure that holds all core dependencies
// and provides methods to manage them. It serves as the central container for
// configuration, database, CLI commands, HTTP server, and validation.
type App struct {
	ConfigLoader *config.Loader
	DB           *gorm.DB
	Command      *cobra.Command
	Gin          *gin.Engine
	Validator    *validator.Validate
}

// SetConfigLoader sets the configuration loader instance for the application.
// This allows you to inject your own configuration loader that implements
// the config contract interface.
func (app *App) SetConfigLoader(loader *config.Loader) {
	app.ConfigLoader = loader
}

// SetValidator sets the validator instance for request validation.
// The validator is used to validate incoming HTTP requests and ensure
// data integrity before processing.
func (app *App) SetValidator(validate *validator.Validate) {
	app.Validator = validate
}

// GetValidator returns the current validator instance used by the application.
// This allows handlers and other components to access the validator for
// performing request validation operations.
func (app *App) GetValidator() *validator.Validate {
	return app.Validator
}

// GetDb returns the GORM database connection instance.
// This provides access to the database for performing CRUD operations
// and executing database queries throughout the application.
func (app *App) GetDb() *gorm.DB {
	return app.DB
}

// GetGin returns the Gin HTTP router engine instance.
// This allows you to register routes, middleware, and handlers
// for your HTTP endpoints.
func (app *App) GetGin() *gin.Engine {
	return app.Gin
}

// SetGin sets the Gin HTTP router engine instance for the application.
// This is typically called by command handlers to configure the HTTP server
// with custom middleware and route setup.
func (app *App) SetGin(engine *gin.Engine) {
	app.Gin = engine
}

// SetupBaseDependencies initializes the core application dependencies.
// This includes setting up the logger, database connection, and other
// essential services required for the application to function.
func (app *App) SetupBaseDependencies() {
	//app.setupConfig()
	app.setupLogger()
	//app.setupDatabaseConnection()
	app.setupModel()
}

// AddCommand registers a new CLI command to the application.
// The provided function should return a CommandContract implementation that
// defines the command's behavior, usage, and execution logic.
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

// Run starts the CLI application and executes the registered commands.
// This method initializes the CLI interface and begins processing
// user commands from the command line.
func (app *App) Run() {
	cliApp := cli.NewCli(app)
	//routers.NewCliRouter(cliApp)
	cliApp.Run()
}

// GetConfigLoader returns the configuration loader instance.
// This provides access to application configuration values loaded from
// environment variables and configuration files.
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

// NewApp creates and returns a new application instance.
// This is the entry point for initializing the Ponodo framework.
// The returned instance implements the AppContract interface.
func NewApp() contracts.AppContract {
	return &App{}
}
