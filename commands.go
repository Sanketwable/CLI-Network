package main

import (
	"time"
	"github.com/spf13/cobra"
)

func printcommand() *cobra.Command {
	return &cobra.Command{
		Use: "curtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("hey sanket current time is", prettyTime)
			return nil
		},
	}
}
