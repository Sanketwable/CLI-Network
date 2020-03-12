package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// wifiCmd represents the wifi command
var wifiCmd = &cobra.Command{
	Use:   "wifi",
	Short: "Used to connect to wifi-board/esp8266",
	Long: `this command is used to connect the board to the wifi using same netwok on PC
	and board to upload code`,
	Example: "CLI-Network network wifi <flags>",
	Run: wififunc,
}

func init() {
	networkCmd.AddCommand(wifiCmd)
	wifiCmd.Flags().BoolP("scan", "s", false, "Scan for the network")
	wifiCmd.Flags().BoolP("connect", "c", false, "Connect to the network")
}
func wififunc(cmd *cobra.Command, args []string) {
	scanFlagStatus, _ := cmd.Flags().GetBool("scan")
	if scanFlagStatus {
		fmt.Println("Wifi scanned")
	} else {
		fmt.Println("Network need subcommands")
	}
}
