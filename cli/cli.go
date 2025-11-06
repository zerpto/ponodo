package cli

import (
	"fmt"
	clicontracts "github.com/zerpto/ponodo/cli/contracts"
	"github.com/zerpto/ponodo/contracts"
	"os"

	"github.com/spf13/cobra"
)

type Cli struct {
	App     contracts.AppContract
	Command *cobra.Command
}

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

func (cli *Cli) SetRootCommand(cmd *cobra.Command) {
	cli.Command = cmd
}

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
