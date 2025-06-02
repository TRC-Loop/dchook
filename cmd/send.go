/*
Copyright © 2025 TRC-Loop <ak@stellar-code.com>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	sendURL     string
	sendMessage string
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to a Discord webhook",
	RunE: func(cmd *cobra.Command, args []string) error {
		if sendURL == "" {
			return fmt.Errorf("please provide --url")
		}
		if sendMessage == "" {
			return fmt.Errorf("please provide --message")
		}

		payload := map[string]string{"content": sendMessage}
		data, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		resp, err := http.Post(sendURL, "application/json", bytes.NewReader(data))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return fmt.Errorf("failed to send message: status %d", resp.StatusCode)
		}

		fmt.Println("Message sent successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(&sendURL, "url", "u", "", "Discord webhook URL")
	sendCmd.Flags().StringVarP(&sendMessage, "message", "m", "", "Message content")
}
