package main

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("Hello")
	cmd := &cobra.Command {
		//Use: "gifm",
		Short: "Hello Gofers",
		SilenceUsage: true,
	}
	cmd.AddCommand(printcommand())

	if err := cmd.Execute(); err != nil {
			os.Exit(1)
		}
}