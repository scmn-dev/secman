package commands

fetch_w := `
	$lastDir = pwd
	cd $HOME/.secman
	git pull
	cd $lastDir
`

fetch_ml := `
	cd ~/.secman
	git pull
	cd -
`

upgrade :=
	`
		l=$(curl --silent "https://api.github.com/repos/secman-team/secman/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
		c=$(secman verx | tr -d '\n')
		smLoc="/usr/local/bin"

		if [ $l == $c ]; then
			echo "secman is already up-to-date and it's the latest release $l"

		elif [ $l != $c ]; then
			sudo rm -rf $smLoc/secman*
			sudo rm -rf $smLoc/cgit*
			sudo rm -rf $smLoc/verx*

			curl -fsSL https://secman-team.github.io/install.sh | bash

			if [ -x "command -v $(secman)" ]; then
				echo "secman was upgraded successfully 🎊"
			fi
		fi
	`

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

push_w := 
	`
		$lastDir = pwd
		cd ~/.secman

		if (Test-Path -path .git) {
			git add .
			git commit -m "new change"
			git push
		}

		cd $lastDir
	`

push_ml := 
	`
		cd ~/.secman
		git add .
		git commit -m "new secman password"
		git push
		cd -
	`

pull_w := 
	`
		$lastDir = pwd
		cd ~/.secman

		git pull

		cd $lastDir
	`

pull_ml :=
	`
		cd ~/.secman
		git pull
		cd -
	`

clone_w := 
	`
		$clone=secman repo clone $SM_GH_UN/.secman ~/.secman

		if (Test-Path -path ~/.secman) {
			Remove-Item ~/.secman -Recurse -Force
			$clone
		} else {
			$clone
		}
	`
clone_ml :=
	`
		clone="secman repo clone $SM_GH_UN/.secman ~/.secman"

		if [ -d ~/.secman ]; then
			rm -rf ~/.secman
			${clone}
		else
			${clone}
		fi
	`

clone_check_w :=
	`
		if (Test-Path -path ~/.secman) {
			Write-Host "cloned successfully"
		}
	`

clone_check_ml := `if [ -d ~/.secman ]; then echo "cloned successfully ✅"; fi`
