# Delete Command

Delete a Discord webhook by its URL.

```bash
dchook delete --url <webhook-url>
```

## Options

| Flag    | Shortcut | Description         | Required |
| ------- | -------- | ------------------- | -------- |
| `--url` | `-u`     | Discord webhook URL | ✅        |

## Example

```bash
dchook delete -u https://discord.com/api/webhooks/...
```

✅ Returns success if the webhook is deleted (`204 No Content`).
❌ Returns an error if the webhook is invalid or already deleted.