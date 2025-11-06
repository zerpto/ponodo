package contracts

import "github.com/spf13/cobra"

// CommandContract defines the interface for CLI command implementations.
// Commands implementing this interface provide usage information, descriptions,
// examples, and execution logic for CLI subcommands.
//
//go:generate mockgen -source=$GOFILE -destination=./mocks/mock_command_contract.go -package=mocks
type CommandContract interface {
	Use() string
	Short() string
	Long() string
	Example() string
	Run(cmd *cobra.Command, args []string)
}
