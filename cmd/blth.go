package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// blthCmd represents the blth command
var blthCmd = &cobra.Command{
	Use:   "blth",
	Short: "Used to connect to H-05 bluetooth module",
	Long: `This command is used to upload code in arduino using
	bluetooth module H-05 over the air without any wire connections`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("blth called")
	},
}

func init() {
	networkCmd.AddCommand(blthCmd)
	blthCmd.Flags().BoolP("scan", "s", false, "Scan for WiFi network")
	blthCmd.Flags().BoolP("connect", "c", false, "Connect to BlueTooth network")

}
