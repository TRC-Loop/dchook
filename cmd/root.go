/*
Copyright © 2025 TRC-Loop <ak@stellar-code.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dchook",
	Short: "A Tool to do all with Discord's Webhooks",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// empty (:
}
