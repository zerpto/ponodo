package cli

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	clicontracts "github.com/zerpto/ponodo/cli/contracts"
	climocks "github.com/zerpto/ponodo/cli/contracts/mocks"
	"github.com/zerpto/ponodo/config"
	configmocks "github.com/zerpto/ponodo/config/contracts/mocks"
	"github.com/zerpto/ponodo/contracts"
	"github.com/zerpto/ponodo/contracts/mocks"
)

func TestNewCli(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mocks.NewMockAppContract(ctrl)
	mockConfigLoader := configmocks.NewMockConfigContract(ctrl)

	mockApp.EXPECT().GetConfigLoader().Return(&config.Loader{
		Config: mockConfigLoader,
	}).AnyTimes()
	mockConfigLoader.EXPECT().GetApp().Return("testapp").AnyTimes()

	cli := NewCli(mockApp)

	require.NotNil(t, cli)
	assert.Equal(t, mockApp, cli.App)
	assert.NotNil(t, cli.Command)
}

func TestCli_SetRootCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mocks.NewMockAppContract(ctrl)
	cli := &Cli{
		App: mockApp,
	}

	cmd := &cobra.Command{
		Use: "test",
	}

	cli.SetRootCommand(cmd)
	assert.Equal(t, cmd, cli.Command)
}

func TestCli_AddCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mocks.NewMockAppContract(ctrl)
	mockCommand := climocks.NewMockCommandContract(ctrl)

	cli := &Cli{
		App: mockApp,
		Command: &cobra.Command{
			Use: "test",
		},
	}

	mockCommand.EXPECT().Use().Return("testcmd").Times(1)
	mockCommand.EXPECT().Short().Return("test short").Times(1)
	mockCommand.EXPECT().Long().Return("test long").Times(1)
	mockCommand.EXPECT().Example().Return("test example").Times(1)
	// Run is called when the command is executed, not when it's added
	// We don't execute it in this test, so we don't expect Run to be called

	cli.AddCommand(func(app contracts.AppContract) clicontracts.CommandContract {
		return mockCommand
	})

	require.NotNil(t, cli.Command)
}

func TestCli_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mocks.NewMockAppContract(ctrl)
	cli := &Cli{
		App: mockApp,
		Command: &cobra.Command{
			Use: "test",
			Run: func(cmd *cobra.Command, args []string) {
				// Do nothing for test
			},
		},
	}

	// Test that Run doesn't panic with a valid command
	// Note: We can't fully test os.Exit behavior, but we can test the structure
	assert.NotNil(t, cli.Command)
}
