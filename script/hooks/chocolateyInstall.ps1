$packageName = 'secman'
$fileType = 'msi'
$url = 'https://github.com/secman-team/secman/releases/download/v5.1.2/secman-windows-v5.1.2.msi'
$silentArgs = '/S'

Install-ChocolateyPackage $packageName $fileType $silentArgs $url
