const LOC = `$loc = "$HOME/AppData/Local/secman"`;

const LastCheck = `
  if (!(Test-Path -path $loc)) {
    Write-Host "secman was uninstalled successfully... thank you for using secman"
  } else {
    Write-Host "there's an error while uninstalling secman, try again"
  }
`;

const MainCode = `
  Remove-Item $loc -Recurse -Force
  $path = $Env:Path
  $newpath = $path.replace("$loc;","")
  $env:Path = $newpath
`;

const CDCmd = `
  Remove-Item $HOME/.secman -Recurse -Force
  $SM_GH_UN = git config user.name

  Write-Host "after clear, if you want to restore .secman you can clone it from your private repo in https://github.com/$SM_GH_UN/.secman"
`;

const Check = `
  $releases = "https://api.github.com/repos/scmn-dev/secman/releases"

  $l = (Invoke-WebRequest -Uri $releases -UseBasicParsing | ConvertFrom-Json)[0].tag_name

  $c = secman verx
`;

module.exports = { LOC, LastCheck, MainCode, CDCmd, Check };
