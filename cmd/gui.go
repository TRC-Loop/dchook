package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "Start the CLI GUI",
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()

		menu := tview.NewList().
			AddItem("Send", "Send a webhook", 's', nil).
			AddItem("Info", "Show webhook info", 'i', nil).
			AddItem("Delete", "Delete a webhook", 'd', nil).
			AddItem("Exit", "Exit the app", 'e', nil)

		menu.SetBorder(true).
			SetTitle("dcHook V" + version).
			SetTitleAlign(tview.AlignCenter)

		mainFlex := tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(
				tview.NewFlex().SetDirection(tview.FlexColumn).
					AddItem(nil, 0, 1, false).
					AddItem(menu, 0, 2, true).
					AddItem(nil, 0, 1, false),
				0, 3, true,
			).
			AddItem(nil, 0, 1, false)

		var sendForm *tview.Form

		createSendWebhookForm := func(app *tview.Application, mainLayout tview.Primitive) *tview.Form {
			sendForm = tview.NewForm().
				AddInputField("Webhook URL:", "", 80, nil, nil).
				AddInputField("Message:", "", 80, nil, nil)

			sendForm.AddButton("Send", func() {
				url := sendForm.GetFormItemByLabel("Webhook URL:").(*tview.InputField).GetText()
				message := sendForm.GetFormItemByLabel("Message:").(*tview.InputField).GetText()

				go func() {
					resultText := ""
					if url == "" {
						resultText = "Error: Please provide a webhook URL."
					} else if message == "" {
						resultText = "Error: Please provide a message."
					} else {
						err := doSendWebhook(url, message)
						if err != nil {
							resultText = fmt.Sprintf("Error sending message:\n%s", err.Error())
						} else {
							resultText = "Message sent successfully!"
							sendForm.GetFormItemByLabel("Webhook URL:").(*tview.InputField).SetText("")
							sendForm.GetFormItemByLabel("Message:").(*tview.InputField).SetText("")
						}
					}

					app.QueueUpdateDraw(func() {
						modal := tview.NewModal().
							SetText(resultText).
							AddButtons([]string{"OK"}).
							SetDoneFunc(func(buttonIndex int, buttonLabel string) {
								app.SetRoot(sendForm, true).SetFocus(sendForm)
							})
						app.SetRoot(modal, true).SetFocus(modal)
					})
				}()
			}).
				AddButton("Back to Menu", func() {
					app.SetRoot(mainLayout, true).SetFocus(menu)
				})

			sendForm.SetBorder(true).SetTitle("Send Webhook").SetTitleAlign(tview.AlignCenter)
			return sendForm
		}

		var deleteForm *tview.Form

		createDeleteWebhookForm := func(app *tview.Application, mainLayout tview.Primitive) *tview.Form {
			deleteForm = tview.NewForm().
				AddInputField("Webhook URL to Delete:", "", 80, nil, nil)

			deleteForm.AddButton("Delete", func() {
				url := deleteForm.GetFormItemByLabel("Webhook URL to Delete:").(*tview.InputField).GetText()

				go func() {
					resultText := ""
					if url == "" {
						resultText = "Error: Please provide a webhook URL to delete."
					} else {
						err := doDeleteWebhook(url)
						if err != nil {
							resultText = fmt.Sprintf("Error deleting webhook:\n%s", err.Error())
						} else {
							resultText = "Webhook deleted successfully!"
							deleteForm.GetFormItemByLabel("Webhook URL to Delete:").(*tview.InputField).SetText("")
						}
					}

					app.QueueUpdateDraw(func() {
						modal := tview.NewModal().
							SetText(resultText).
							AddButtons([]string{"OK"}).
							SetDoneFunc(func(buttonIndex int, buttonLabel string) {
								app.SetRoot(deleteForm, true).SetFocus(deleteForm)
							})
						app.SetRoot(modal, true).SetFocus(modal)
					})
				}()
			}).
				AddButton("Back to Menu", func() {
					app.SetRoot(mainLayout, true).SetFocus(menu)
				})

			deleteForm.SetBorder(true).SetTitle("Delete Webhook").SetTitleAlign(
				tview.AlignCenter,
			)
			return deleteForm
		}

		var infoForm *tview.Form

		createInfoWebhookForm := func(app *tview.Application, mainLayout tview.Primitive) *tview.Form {
			infoForm = tview.NewForm().
				AddInputField("Webhook URL for Info:", "", 80, nil, nil)

			infoForm.AddButton("Get Info", func() {
				url := infoForm.GetFormItemByLabel("Webhook URL for Info:").(*tview.InputField).GetText()

				go func() {
					infoText := ""
					formattedInfo, rawJSONBytes, err := doGetWebhookInfo(url)
					if err != nil {
						infoText = fmt.Sprintf("Error getting webhook info:\n%s", err.Error())
					} else {
						infoDisplayView := tview.NewTextView()
						infoDisplayView.SetDynamicColors(true)
						infoDisplayView.SetRegions(true)
						infoDisplayView.SetWrap(true)
						infoDisplayView.SetBorder(true)
						infoDisplayView.SetTitle("Webhook Info")
						infoDisplayView.SetTitleAlign(tview.AlignCenter)

						renderContent := func(isRaw bool) {
							infoDisplayView.Clear()
							if isRaw {
								infoDisplayView.SetText(string(rawJSONBytes))
							} else {
								infoDisplayView.SetText(formattedInfo)
							}
						}

						infoDisplayView.SetInputCapture(
							func(event *tcell.EventKey) *tcell.EventKey {
								if event.Key() == tcell.KeyRune && event.Rune() == 'r' {
									currentText := infoDisplayView.GetText(true)
									isRawCurrently := (currentText == string(rawJSONBytes))
									renderContent(!isRawCurrently)
									return nil
								}
								if event.Key() == tcell.KeyEscape {
									app.SetRoot(infoForm, true).SetFocus(infoForm)
									return nil
								}
								return event
							},
						)

						app.QueueUpdateDraw(func() {
							renderContent(false)
							infoDisplayView.SetBorder(true).SetTitle("Press 'r' for Raw JSON / 'Esc' to go back")
							app.SetRoot(infoDisplayView, true).SetFocus(infoDisplayView)
						})
					}

					if infoText != "" {
						app.QueueUpdateDraw(func() {
							modal := tview.NewModal().
								SetText(infoText).
								AddButtons([]string{"OK"}).
								SetDoneFunc(func(buttonIndex int, buttonLabel string) {
									app.SetRoot(infoForm, true).SetFocus(infoForm)
								})
							app.SetRoot(modal, true).SetFocus(modal)
						})
					}
				}()
			}).
				AddButton("Back to Menu", func() {
					app.SetRoot(mainLayout, true).SetFocus(menu)
				})

			infoForm.SetBorder(true).SetTitle("Get Webhook Info").SetTitleAlign(
				tview.AlignCenter,
			)
			return infoForm
		}

		menu.SetSelectedFunc(func(index int, mainText, secondaryText string,
			shortcut rune) {
			switch mainText {
			case "Send":
				sendForm = createSendWebhookForm(app, mainFlex)
				app.SetRoot(sendForm, true).SetFocus(sendForm)
			case "Info":
				infoForm = createInfoWebhookForm(app, mainFlex)
				app.SetRoot(infoForm, true).SetFocus(infoForm)
			case "Delete":
				deleteForm = createDeleteWebhookForm(app, mainFlex)
				app.SetRoot(deleteForm, true).SetFocus(deleteForm)
			case "Exit":
				app.Stop()
			}
		})

		if err := app.SetRoot(mainFlex, true).SetFocus(menu).Run(); err != nil {
			panic(err)
		}
	},
}

func doSendWebhook(url, message string) error {
	payload := map[string]string{"content": message}
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON payload: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		responseBody := new(bytes.Buffer)
		responseBody.ReadFrom(resp.Body)
		return fmt.Errorf(
			"failed to send message, status code: %d, response: %s",
			resp.StatusCode,
			responseBody.String(),
		)
	}

	return nil
}

func doDeleteWebhook(url string) error {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create DELETE request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		responseBody := new(bytes.Buffer)
		responseBody.ReadFrom(resp.Body)
		return fmt.Errorf(
			"failed to delete webhook, status code: %d, response: %s",
			resp.StatusCode,
			responseBody.String(),
		)
	}

	return nil
}

func doGetWebhookInfo(url string) (string, []byte, error) {
	type WebhookInfoInternal struct {
		ID        string  `json:"id"`
		Type      int     `json:"type"`
		GuildID   *string `json:"guild_id"`
		ChannelID string  `json:"channel_id"`
		Name      *string `json:"name"`
		Avatar    *string `json:"avatar"`
		Token     string  `json:"token"`
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != 200 {
		return "", body, fmt.Errorf(
			"failed to get info, status code: %d, response: %s",
			resp.StatusCode,
			string(body),
		)
	}

	var webhook WebhookInfoInternal
	if err := json.Unmarshal(body, &webhook); err != nil {
		return "", body, fmt.Errorf("failed to parse webhook info: %w", err)
	}

	safeString := func(s *string) string {
		if s == nil || strings.TrimSpace(*s) == "" {
			return "N/A"
		}
		return *s
	}

	formatAvatar := func(a *string) string {
		if a == nil || *a == "" {
			return "N/A"
		}
		return *a
	}

	var builder strings.Builder
	fmt.Fprintf(&builder, "[blue]Webhook ID: [white]%s\n", webhook.ID)
	fmt.Fprintf(&builder, "[blue]Type: [white]%d\n", webhook.Type)
	fmt.Fprintf(&builder, "[blue]Guild ID: [white]%s\n", safeString(webhook.GuildID))
	fmt.Fprintf(&builder, "[blue]Channel ID: [white]%s\n", webhook.ChannelID)
	fmt.Fprintf(&builder, "[blue]Name: [white]%s\n", safeString(webhook.Name))
	fmt.Fprintf(&builder, "[blue]Avatar: [white]%s\n", formatAvatar(webhook.Avatar))
	fmt.Fprintf(&builder, "[blue]Token: [white]%s\n", webhook.Token)

	return builder.String(), body, nil
}

func init() {
	rootCmd.AddCommand(guiCmd)
}
