/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func timefunc(cmd *cobra.Command, args []string) error {
	fstatus, _ := cmd.Flags().GetBool("ruby")
	
	if !fstatus {
		now := time.Now()
		fmt.Println(now)
	} else {
		now := time.Now().UTC()
		fmt.Println(now)
	}
	return nil
}