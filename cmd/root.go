package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "CLI-Network",
	Short: "A Network discovery CLI",
	Long: `It is a CLI-Interface for discovery of all kinds of devices and boards via
	1. WIFI
	2. Bluetooth.`,
	Example:          "  " + os.Args[0] + " <command> [flags...]",
	PersistentPreRun: rootfunction,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func rootfunction(cmd *cobra.Command, args []string) {
		cmd.Println("started CLI")
}

