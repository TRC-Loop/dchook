# dchook

A simple CLI tool to manage Discord webhooks: info, delete and send messages.

<p align="center">
  <img src="https://github.com/TRC-Loop/dchook/blob/main/.github/assets/dcHook.svg" alt="dcHook Logo" width="350" />
</p>

<p align="center">
  <img alt="GitHub Tag" src="https://img.shields.io/github/v/tag/TRC-Loop/dchook?style=for-the-badge&label=Version">
  <img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/TRC-Loop/dchook?style=for-the-badge">
  <img alt="GitHub forks" src="https://img.shields.io/github/forks/TRC-Loop/dchook?style=for-the-badge">
  <img alt="GitHub License" src="https://img.shields.io/github/license/TRC-Loop/dchook?style=for-the-badge">
</p>

<div align="center">

<table>
  <tr>
    <td align="center">
      <strong>Stack</strong><br>
      <img alt="Skill Icons" src="https://skillicons.dev/icons?i=go,discord&perline=42&theme=dark">
    </td>
    <td align="center">
      <strong>Runs on</strong><br>
      <img alt="Skill Icons" src="https://skillicons.dev/icons?i=windows,apple,linux&perline=42&theme=dark">
    </td>
  </tr>
</table>

</div>




---

## Features

- Get webhook info
- Delete a webhook
- Send a message to a webhook

---

## Requirements

- Go 1.19+ installed
- Internet connection

---

## Installation

1. Clone the repo:

```bash
git clone https://github.com/TRC-Loop/dchook.git
cd dchook
````

2. Build the executable:

**Linux, MacOS**

```bash
go build -o dchook
```

**Windows**

```bash
go build -ldflags="-s -w" -o dchook.exe
```
> **Note:** ldflages are used because Windows Defender detects this as a False-Positive.

3. (Optional) Add `dchook` to your PATH to use it anywhere:

**Linux, MacOS**

* Move the executable to `/usr/local/bin` (or any folder in your PATH):

```bash
sudo mv dchook /usr/local/bin/
```

* Now you can run:

```bash
dchook --help
```

**Windows**

* Move `dchook.exe` to a folder, for example `C:\Tools`

* Add that folder to your PATH:

  1. Open Start Menu > Search "Environment Variables" > Edit system environment variables
  2. Click "Environment Variables"
  3. Select "Path" under your system variables, then "Edit"
  4. Click "New" and add `C:\Tools\` (or your folder)
  5. Click OK on all dialogs
  6. Restart your terminal or PowerShell
	> **Note:** If it's still not working, restart your PC.

* Now you can run:

```powershell
dchook.exe --help
```

---

## Usage examples

**Get webhook info:**

```bash
dchook info --url <webhook_url>
```

> **Note:** use `--raw` to get the Raw, Non-formatted JSON

**Delete a webhook:**

```bash
dchook delete --url <webhook_url>
```

**Send a message:**

```bash
dchook send --url <webhook_url> -m "Hello from dchook!"
```

---

## Notes

* I am not responsible for any damage done with this Software.
* Make sure the webhook token is valid and you have permission to manage it.

---

## ToDo

* Command Line GUI (`dchook gui`)
* Actual GUI?
* More Options to save to file
* More Debug Info (`--debug`)
* Copy Info to clipboard (specific Info/all Info `--copy`)
* Good Docs (McDocs?)
* Installer/Installer Script
* On Release: Build Executable for Linux/Mac

---

## License

MIT License
