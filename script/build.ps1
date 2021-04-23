$LOC = "$HOME\AppData\Local\secman"

if ((Get-Command ruby -errorAction SilentlyContinue) -or (Get-Command git -errorAction SilentlyContinue)) {
    $LATEST_VERSION=git describe --abbrev=0 --tags
    $VCURL = "https://raw.githubusercontent.com/secman-team/tools/HEAD/sm.ps1"
    $BIN = "$LOC\bin"

    # Build
    cd core
    go get -d
    go build -o secman.exe -ldflags "-X main.version=$LATEST_VERSION"
    git clone https://github.com/secman-team/sm-win $HOME/sm
    Invoke-WebRequest $VCURL -outfile $HOME\sm\sm.ps1

    # Setup
    gem install colorize
    New-Item -ItemType "directory" -Path $BIN
    Move-Item secman.exe -Destination $BIN
    [System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$BIN", [System.EnvironmentVariableTarget]::User)

    cd ..
} else {
    Write-Host "Some of these apps must be installed: git, or ruby"
    Write-Host "git: https://git-scm.com"
    Write-Host "ruby: https://rubyinstaller.org"
}

if (Test-Path -path $LOC) {
    Write-Host "secman was builded successfully, run secman --help" -ForegroundColor DarkGreen
} else {
    Write-Host "Build failed"
}
