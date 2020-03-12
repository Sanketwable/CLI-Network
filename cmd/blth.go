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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blthCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
