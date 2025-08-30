package was

import (
	"github.com/spf13/cobra"
)

var (
	Port = 8080
)

var brokerCmd = &cobra.Command{
	Use:   "broker",
	Short: "Manage WebAssembly brokers",
	Long:  `A command for managing WebAssembly brokers.`,
	Run: func(cmd *cobra.Command, args []string) {
		IsBroker = true
	},
}

func init() {
	rootCmd.AddCommand(brokerCmd)
	brokerCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8080, "Broker port")
}
