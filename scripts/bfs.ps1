# Build From Source
$loc = "$HOME\AppData\Local\secman"

go run scripts/date.go >> date.txt

$LATEST_VERSION=git describe --abbrev=0 --tags
$DATE=cat date.txt

# Build
go mod tidy -compat="1.18"
go build -o secman.exe -ldflags "-X main.version=$LATEST_VERSION -X main.versionDate=$DATE"

# Setup
$BIN = "$loc\bin"
New-Item -ItemType "directory" -Path $BIN
Move-Item secman.exe -Destination $BIN
[System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$BIN", [System.EnvironmentVariableTarget]::User)

if (Test-Path -path $loc) {
    Write-Host "Secman was built successfully, refresh your powershell and then run 'secman --help'" -ForegroundColor DarkGreen
} else {
    Write-Host "Build failed" -ForegroundColor Red
}
