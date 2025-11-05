package contracts

import "github.com/spf13/cobra"

type CommandContract interface {
	Use() string
	Short() string
	Long() string
	Example() string
	Run(cmd *cobra.Command, args []string)
}
