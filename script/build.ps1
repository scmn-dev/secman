# Build From Source
$loc = "$HOME\AppData\Local\secman"

if ((Get-Command git -errorAction SilentlyContinue) -or (Get-Command npm -errorAction SilentlyContinue)) {
    $LATEST_VERSION=git describe --abbrev=0 --tags
    $DATE=git tag -l --sort=-creatordate --format='%(creatordate:short)' $LATEST_VERSION

    # Build
    cd core
    go get -d
    go build -o secman.exe -ldflags "-X main.version=$LATEST_VERSION -X main.versionDate=($DATE)"

    # Setup
    $BIN = "$loc\bin"
    New-Item -ItemType "directory" -Path $BIN
    Move-Item secman.exe -Destination $BIN
    [System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$BIN", [System.EnvironmentVariableTarget]::User)

} else {
    Write-Host "Some of these apps must be installed: git, or npm"
    Write-Host "git: https://git-scm.com"
    Write-Host "npm: https://nodejs.org"
}

if (Test-Path -path $loc) {
    Write-Host "Secman was built successfully, run `secman --help` " -ForegroundColor DarkGreen
} else {
    Write-Host "Build failed"
}
