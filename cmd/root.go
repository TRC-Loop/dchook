/*
Copyright © 2025 TRC-Loop <ak@stellar-code.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dchook",
	Short: "A Tool to do all with Discord's Webhooks",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// empty (:
}
