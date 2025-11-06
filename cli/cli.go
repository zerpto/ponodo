package cli

import (
	"fmt"
	"os"

	clicontracts "github.com/zerpto/ponodo/cli/contracts"
	"github.com/zerpto/ponodo/contracts"

	"github.com/spf13/cobra"
)

// Cli represents the command-line interface structure that manages
// CLI commands and their execution. It wraps the Cobra command framework
// and provides a clean interface for registering and running commands.
type Cli struct {
	App     contracts.AppContract
	Command *cobra.Command
}

// Run executes the CLI application by processing the root command.
// It handles command execution and exits with an error code if
// the command fails or encounters an error.
func (cli *Cli) Run() {
	rootCmd := cli.Command
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

// SetRootCommand sets the root Cobra command for the CLI application.
// This command serves as the entry point for all registered subcommands
// and defines the base command structure.
func (cli *Cli) SetRootCommand(cmd *cobra.Command) {
	cli.Command = cmd
}

// AddCommand registers a new subcommand to the CLI application.
// The provided function should return a CommandContract implementation
// that defines the command's behavior, usage, and execution logic.
func (cli *Cli) AddCommand(f func(app contracts.AppContract) clicontracts.CommandContract) {
	rootCmd := cli.Command

	command := f(cli.App)

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

// NewCli creates and initializes a new CLI application instance.
// It sets up the root command based on the application configuration
// and returns a ready-to-use CLI instance.
func NewCli(app contracts.AppContract) *Cli {
	cli := &Cli{
		App: app,
	}

	config := cli.App.GetConfigLoader().Config
	var rootCmd = &cobra.Command{
		Use:   config.GetApp(),
		Short: fmt.Sprintf("%s Service", config.GetApp()),
	}
	cli.SetRootCommand(rootCmd)
	return cli
}
