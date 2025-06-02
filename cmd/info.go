/*
Copyright © 2025 TRC Loop <ak@stellar-code.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var (
	infoURL string
	rawJSON bool
	pretty  bool
)

type WebhookInfo struct {
	ID        string  `json:"id"`
	Type      int     `json:"type"`
	GuildID   *string `json:"guild_id"`
	ChannelID string  `json:"channel_id"`
	Name      *string `json:"name"`
	Avatar    *string `json:"avatar"`
	Token     string  `json:"token"`
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about a Discord webhook",
	RunE: func(cmd *cobra.Command, args []string) error {
		if infoURL == "" {
			return fmt.Errorf("please provide --url")
		}

		resp, err := http.Get(infoURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("failed to get info: status %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if rawJSON {
			// Raw output
			fmt.Println(string(body))
			return nil
		}

		var webhook WebhookInfo
		if err := json.Unmarshal(body, &webhook); err != nil {
			if pretty {
				// fallback: pretty print raw JSON
				var rawObj map[string]interface{}
				if err2 := json.Unmarshal(body, &rawObj); err2 == nil {
					out, _ := json.MarshalIndent(rawObj, "", "  ")
					fmt.Println(string(out))
					return nil
				}
			}
			// fallback raw print if parsing fails
			fmt.Println(string(body))
			return nil
		}

		// Default: nicely formatted human readable
		fmt.Printf("Webhook ID: %s\n", webhook.ID)
		fmt.Printf("Type: %d\n", webhook.Type)
		fmt.Printf("Guild ID: %s\n", safeString(webhook.GuildID))
		fmt.Printf("Channel ID: %s\n", webhook.ChannelID)
		fmt.Printf("Name: %s\n", safeString(webhook.Name))
		fmt.Printf("Avatar: %s\n", formatAvatar(webhook.Avatar))
		fmt.Printf("Token: %s\n", webhook.Token)

		return nil
	},
}

func safeString(s *string) string {
	if s == nil || strings.TrimSpace(*s) == "" {
		return "N/A"
	}
	return *s
}

func formatAvatar(a *string) string {
	if a == nil || *a == "" {
		return "N/A"
	}
	return *a
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().StringVarP(&infoURL, "url", "u", "", "Discord webhook URL")
	infoCmd.Flags().BoolVarP(&rawJSON, "raw", "r", false, "Print raw JSON")
}
