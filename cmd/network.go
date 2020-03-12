package cmd

import (
	"github.com/spf13/cobra"
)

// networkCmd represents the network command
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "It is used to view and connect network available",
	Long: `This is a command used to search for bluetooth and wifi boards near and connect them to upload code OTA`,
	Example: "CLI-Network network [wifi/blth] <flags>",
}

func init() {
	rootCmd.AddCommand(networkCmd)
}
