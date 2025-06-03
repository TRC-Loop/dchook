function Start-Spinner {
    param([string]$Message)
    $Global:spinnerActive = $true
    $spinner = "/-|"
    $i = 0
    Write-Host -NoNewline "$Message "
    while ($Global:spinnerActive) {
        Write-Host -NoNewline "b$($spinner[$i % $spinner.Length])"
        Start-Sleep -Milliseconds 100
        $i++
    }
    Write-Host "b "
}

function Stop-Spinner {
    $Global:spinnerActive = $false
}

Write-Host "[+] Cloning dchook repo..."
Start-Job -ScriptBlock {
    git clone https://github.com/TRC-Loop/dchook.git
} | Out-Null
Start-Spinner "Cloning..."
Wait-Job * | Out-Null
Stop-Spinner
Set-Location dchook

Write-Host "[+] Building dchook.exe..."
Start-Job -ScriptBlock {
    go build -ldflags="-s -w" -o dchook.exe
} | Out-Null
Start-Spinner "Building..."
Wait-Job * | Out-Null
Stop-Spinner

if (-Not (Test-Path ".\dchook.exe")) {
    Write-Host "[-] Build failed. Make sure Go is installed." -ForegroundColor Red
    exit 1
}

$toolsPath = "C:\Tools"
if (-Not (Test-Path $toolsPath)) {
    New-Item -Path $toolsPath -ItemType Directory | Out-Null
}
Copy-Item .\dchook.exe "$toolsPath\dchook.exe" -Force

#Add to path
$envPath = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine)
if ($envPath -notlike "$toolsPath") {
    Write-Host "[+] Adding $toolsPath to system PATH..."
    [Environment]::SetEnvironmentVariable("Path", "$envPath;$toolsPath", [EnvironmentVariableTarget]::Machine)
    Write-Host "[!] You must restart your terminal or PC for changes to take effect." -ForegroundColor Yellow
} else {
    Write-Host "[i] $toolsPath already in PATH"
}

Write-Host "[✓] Done! Run 'dchook.exe --help'"
