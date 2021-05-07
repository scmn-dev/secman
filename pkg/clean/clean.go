package clean

import "github.com/secman-team/shell"

clean_w :=
	`
		$directoyPath = "~/.secman"

		if (Test-Path -path $directoyPath) {
			Remove-Item $directoyPath -Recurse -Force
		}

		if (!(Test-Path -path $directoyPath)) {
			Write-Host "secman was cleaned successfully 🧹"
		}
	`

clean_ml := 
	`
		if [ -d ~/.secman ]; then rm -rf ~/.secman; fi
		if ! [ -d ~/.secman ]; echo "secman was cleaned successfully 🧹"; fi
	`

func Clean() {
	shell.SHCore(clean_ml, clean_w)
}
