package was

import "github.com/spf13/cobra"

var (
	IsAgent bool
	IsBroker bool
)

var rootCmd = &cobra.Command{
	Use:   "was",
	Short: "WebAssembly Command Line Interface",
	Long:  `A command line interface for managing WebAssembly modules.`,
}

func Execute() error {
	return rootCmd.Execute()
}