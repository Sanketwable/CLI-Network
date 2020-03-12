package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"time"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Gives current time of the system",
	Long: `It is a command that provides current time of the system`,
	Example: "#All the commands for time \n    CLI-Network time -now or (-n)\n    CLI-Network time -ruby or (-r)",
	RunE: timefunc,
}

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.Flags().BoolP("ruby", "r", false, "Current Time of the System in ruby format")
	timeCmd.Flags().BoolP("now", "n", false, "Current Time of the System in normal format")
}

func timefunc(cmd *cobra.Command, args []string) error {
	rubyFlagStatus, _ := cmd.Flags().GetBool("ruby")
	nowFlagStatus, _ := cmd.Flags().GetBool("now")

	if nowFlagStatus {
		now := time.Now()
		fmt.Println("Now time is :")
		fmt.Println(now)
	} else if rubyFlagStatus {
		now := time.Now()
		fmt.Println("Time in UTC format is :")
		fmt.Println(now.Format("2006-01-02 15:04:05 Monday"))
	} else {
		now := time.Now()
		fmt.Println("Time is :")
		fmt.Println(now)
	}
	return nil
}
