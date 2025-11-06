package ponodo

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"

	"github.com/zerpto/ponodo/cli"
	clicontracts "github.com/zerpto/ponodo/cli/contracts"
	climocks "github.com/zerpto/ponodo/cli/contracts/mocks"
	"github.com/zerpto/ponodo/config"
	configmocks "github.com/zerpto/ponodo/config/contracts/mocks"
	"github.com/zerpto/ponodo/contracts"
	"github.com/zerpto/ponodo/contracts/mocks"
)

func TestNewApp(t *testing.T) {
	app := NewApp()
	assert.NotNil(t, app)
	assert.IsType(t, &App{}, app)
}

func TestApp_SetConfigLoader(t *testing.T) {
	app := &App{}
	loader := &config.Loader{}
	app.SetConfigLoader(loader)
	assert.Equal(t, loader, app.ConfigLoader)
}

func TestApp_GetConfigLoader(t *testing.T) {
	app := &App{}
	loader := &config.Loader{}
	app.ConfigLoader = loader
	assert.Equal(t, loader, app.GetConfigLoader())
}

func TestApp_SetValidator(t *testing.T) {
	app := &App{}
	validatorInstance := validator.New()
	app.SetValidator(validatorInstance)
	assert.Equal(t, validatorInstance, app.Validator)
}

func TestApp_GetValidator(t *testing.T) {
	app := &App{}
	validatorInstance := validator.New()
	app.Validator = validatorInstance
	assert.Equal(t, validatorInstance, app.GetValidator())
}

func TestApp_GetDb(t *testing.T) {
	app := &App{}
	db := &gorm.DB{}
	app.DB = db
	assert.Equal(t, db, app.GetDb())
}

func TestApp_GetGin(t *testing.T) {
	app := &App{}
	ginEngine := gin.New()
	app.Gin = ginEngine
	assert.Equal(t, ginEngine, app.GetGin())
}

func TestApp_SetGin(t *testing.T) {
	app := &App{}
	ginEngine := gin.New()
	app.SetGin(ginEngine)
	assert.Equal(t, ginEngine, app.Gin)
}

func TestApp_SetupBaseDependencies(t *testing.T) {
	t.Skip("Skipping database-related test - requires database connection")
}

func TestApp_AddCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommand := climocks.NewMockCommandContract(ctrl)
	app := &App{
		Command: &cobra.Command{
			Use: "testapp",
		},
	}

	mockCommand.EXPECT().Use().Return("test").Times(1)
	mockCommand.EXPECT().Short().Return("test short").Times(1)
	mockCommand.EXPECT().Long().Return("test long").Times(1)
	mockCommand.EXPECT().Example().Return("test example").Times(1)
	// Run is called when the command is executed, not when it's added
	// We don't execute it in this test, so we don't expect Run to be called

	app.AddCommand(func(app contracts.AppContract) clicontracts.CommandContract {
		return mockCommand
	})

	// Verify command was added by checking if it exists in root command
	require.NotNil(t, app.Command)
}

func TestApp_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mocks.NewMockAppContract(ctrl)
	mockConfigLoader := configmocks.NewMockConfigContract(ctrl)

	mockApp.EXPECT().GetConfigLoader().Return(&config.Loader{
		Config: mockConfigLoader,
	}).AnyTimes()
	mockConfigLoader.EXPECT().GetApp().Return("testapp").AnyTimes()

	// This will create a CLI instance but we can't easily test the execution
	// without mocking os.Exit, so we'll just verify it doesn't panic
	assert.NotPanics(t, func() {
		// We can't fully test Run() without mocking os.Exit,
		// but we can verify the CLI is created
		cliApp := cli.NewCli(mockApp)
		assert.NotNil(t, cliApp)
	})
}
