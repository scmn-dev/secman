$releases = "https://api.github.com/repos/secman-team/secman/releases"

$tag = (Invoke-WebRequest -Uri $releases -UseBasicParsing | ConvertFrom-Json)[0].tag_name

Write-Host $tag
