/*
Copyright © 2025 TRC Loop <ak@stellar-code.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var version = "1.0.1-01"

var offline bool
var github = "TRC-Loop/dchook"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show dchook version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dchook version: V" + version)

		if offline {
			return
		}

		latest, err := fetchLatestGitHubRelease(github)
		if err == nil {
			fmt.Println("latest on GitHub:", latest)
		}
	},
}

func init() {
	versionCmd.Flags().BoolVarP(&offline, "offline", "o", false, "Disable GitHub check")
	rootCmd.AddCommand(versionCmd)
}

func fetchLatestGitHubRelease(repo string) (string, error) {
	client := &http.Client{Timeout: 3 * time.Second}
	url := "https://api.github.com/repos/" + repo + "/releases/latest"

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	var data struct {
		TagName string `json:"tag_name"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	return data.TagName, nil
}
