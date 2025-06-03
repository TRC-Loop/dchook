#!/bin/bash

spinner() {
    local pid=$!
    local delay=0.1
    local spinstr='|/-'
    while kill -0 $pid 2>/dev/null; do
        local temp=${spinstr#?}
        printf " [%c]  " "$spinstr"
        local spinstr=$temp${spinstr%"$temp"}
        sleep $delay
        printf "\b\b\b\b\b\b"
    done
    printf "    \b\b\b\b"
}

echo "[+] Cloning dchook..."
(git clone https://github.com/TRC-Loop/dchook.git) & spinner
cd dchook  { echo "cd failed."; exit 1; }

echo "[+] Building dchook..."
(go build -o dchook) & spinner
if [ ! -f "./dchook" ]; then
    echo "[-] Build failed. Make sure Go is installed."
    exit 1
fi

read -p "[?] Add dchook to /usr/local/bin so it's globally accessible? (y/n): " add_to_path
if [[ $add_to_path == "y"  $add_to_path == "Y" ]]; then
    echo "[+] Moving dchook to /usr/local/bin..."
    (sudo mv dchook /usr/local/bin/) & spinner
    echo "[✓] dchook installed to /usr/local/bin"
else
    echo "[i] dchook built in ./dchook. Move it manually if needed."
fi

echo "[✓] Done! Try running 'dchook --help'"
