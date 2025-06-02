/*
Copyright © 2025 TRC-Loop <ak@stellar-code.com>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var deleteURL string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Discord webhook",
	RunE: func(cmd *cobra.Command, args []string) error {
		if deleteURL == "" {
			return fmt.Errorf("please provide --url")
		}

		req, err := http.NewRequest("DELETE", deleteURL, nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 204 {
			return fmt.Errorf("failed to delete webhook: status %d", resp.StatusCode)
		}

		fmt.Println("Webhook deleted successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&deleteURL, "url", "u", "", "Discord webhook URL")
}
