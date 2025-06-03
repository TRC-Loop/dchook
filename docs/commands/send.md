# Send Command

Send a message to a Discord webhook.

```bash
dchook send --url <webhook-url> --message "Hello from CLI"
```

## Options

| Flag        | Shortcut | Description             | Required |
| ----------- | -------- | ----------------------- | -------- |
| `--url`     | `-u`     | Discord webhook URL     | ✅        |
| `--message` | `-m`     | Message content to send | ✅        |

## Example

```bash
dchook send -u https://discord.com/api/webhooks/... -m "Hello from dcHook!"
```
