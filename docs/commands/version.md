# Version Command

Show the current version of **dcHook** and optionally check the latest release on GitHub.

```bash
dchook version
```

## Options

| Flag        | Shortcut | Description               | Required |
| ----------- | -------- | ------------------------- | -------- |
| `--offline` | `-o`     | Skip GitHub version check | ❌        |

## Examples

### Show local version and check GitHub:

```bash
dchook version
```

### Show local version only (offline mode):

This won't get the latest release version from Github.

```bash
dchook version --offline
```