# Info Command

Fetch and display information about a Discord webhook.

```bash
dchook info --url <webhook-url>
```

## Options

| Flag    | Shortcut | Description         | Required |
| ------- | -------- | ------------------- | -------- |
| `--url` | `-u`     | Discord webhook URL | ✅        |
| `--raw` | `-r`     | Output raw JSON     | ❌        |

## Output Modes

* **Default**: Clean, human-readable format
* **--raw**: Raw JSON response
* **Fallback**: Pretty JSON if decoding fails and `--raw` is not used

## Examples

### Normal output:

```bash
dchook info -u https://discord.com/api/webhooks/...
```

### Raw JSON:

```bash
dchook info -u https://discord.com/api/webhooks/... --raw
```