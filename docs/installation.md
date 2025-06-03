# Installation

## Requirements

- [Go 1.19+](https://go.dev/doc/install)
- Internet connection

---

## 1. Clone the Repository

```bash
git clone https://github.com/TRC-Loop/dchook.git
cd dchook
```

---

## 2. Build the Executable

### Linux / macOS

```bash
go build -o dchook
```

### Windows

```bash
go build -ldflags="-s -w" -o dchook.exe
```

> ⚠️ `-ldflags="-s -w"` helps avoid false-positives from Windows Defender.

---

## 3. (Optional) Add to PATH

### Linux / macOS

```bash
sudo mv dchook /usr/local/bin/
```

Now run:

```bash
dchook --help
```

### Windows

1. Move `dchook.exe` to a folder like `C:\Tools\`
2. Add that folder to your **System PATH**:

   * Start Menu → Search `Environment Variables`
   * Click **Environment Variables**
   * Edit **Path** → Add `C:\Tools\`
3. Restart PowerShell or Terminal

Then run:

```powershell
dchook.exe --help
```

---

✅ You’re ready to use `dchook`.