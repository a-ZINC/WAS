package was

import (
	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Manage WebAssembly agents",
	Long:  `A command for managing WebAssembly agents.`,
	Run: func(cmd *cobra.Command, args []string) {
		IsAgent = true
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)
}
