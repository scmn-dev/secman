package commands

func Fetch_w() string {
	return `
		$lastDir = pwd
		cd $HOME/.secman
		git pull
		cd $lastDir
	`
} 

func Fetch_ml() string {
	return `
		cd ~/.secman
		git pull
		cd -
	`
}

func Upgrade() string {
	return `
		l=$(curl --silent "https://api.github.com/repos/secman-team/secman/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
		c=$(secman verx | tr -d '\n')
		smLoc="/usr/local/bin"

		if [ $l == $c ]; then
			echo "secman is already up-to-date and it's the latest release $l"

		elif [ $l != $c ]; then
			sudo rm -rf $smLoc/secman*
			sudo rm -rf $smLoc/cgit*
			sudo rm -rf $smLoc/verx*

			curl -fsSL https://deps.secman.dev/install.sh | bash

			if [ -x "command -v $(secman)" ]; then
				echo "secman was upgraded successfully 🎊"
			fi
		fi
	`
}

func Clean_w() string {
	return `
		$directoyPath = "~/.secman"

		if (Test-Path -path $directoyPath) {
			Remove-Item $directoyPath -Recurse -Force
		}

		if (!(Test-Path -path $directoyPath)) {
			Write-Host "secman was cleaned successfully 🧹"
		}
	`
}

func Clean_ml() string {
	return `
		if [ -d ~/.secman ]; then rm -rf ~/.secman; fi
		if ! [ -d ~/.secman ]; echo "secman was cleaned successfully 🧹"; fi
	`
}

func Start_w() string {
	return `
		$SM_GH_UN = git config user.name
		cd ~/.secman

		git init

		echo "# My secman passwords - $SM_GH_UN" >> ~/.secman\README.md

		secman repo create $SM_GH_UN/.secman -y --private

		git add .
		git commit -m "new .secman repo"
		git branch -M trunk
		git remote add origin https://github.com/$SM_GH_UN/.secman
		git push -u origin trunk

		cd $lastDir
	`
}

func Start_ml() string {
	return `
		SM_GH_UN=$(git config user.name)
		cd ~/.secman
		git init

		echo "# My secman passwords - $SM_GH_UN" >> ~/.secman/README.md

		secman repo create $SM_GH_UN/.secman -y --private

		git add .
		git commit -m "new .secman repo"
		git branch -M trunk
		git remote add origin https://github.com/$SM_GH_UN/.secman
		git push -u origin trunk
		cd -
	`
}

func Push_w() string {
	return `
		$lastDir = pwd
		cd ~/.secman

		if (Test-Path -path .git) {
			git add .
			git commit -m "new change"
			git push
		}

		cd $lastDir
	`
}

func Push_ml() string {
	return `
		cd ~/.secman
		git add .
		git commit -m "new secman password"
		git push
		cd -
	`
}

func Pull_w() string {
	return `
		$lastDir = pwd
		cd ~/.secman
		
		git pull
		
		cd $lastDir
	`
}

func Pull_ml() string {
	return `
		cd ~/.secman
		git pull
		cd -
	`
}

func Clone_w() string {
	return `
		$SM_GH_UN = git config user.name
		$clone=secman repo clone $SM_GH_UN/.secman ~/.secman
		
		if (Test-Path -path ~/.secman) {
			Remove-Item ~/.secman -Recurse -Force
			$clone
		} else {
			$clone
		}
	`
}
		
func Clone_ml() string {
	return `
		SM_GH_UN=$(git config user.name)
		clone="secman repo clone $SM_GH_UN/.secman ~/.secman"
		
		if [ -d ~/.secman ]; then
			rm -rf ~/.secman
			${clone}
		else
			${clone}
		fi
	`
}

func Clone_check_w() string {
	return `
		if (Test-Path -path ~/.secman) {
			Write-Host "cloned successfully"
		}
	`
}

func Clone_check_ml() string {
	return `if [ -d ~/.secman ]; then echo "cloned successfully ✅"; fi`
}

func Check_w() string {
	return `
		$releases = "https://api.github.com/repos/secman-team/secman/releases"

		$l = (Invoke-WebRequest -Uri $releases -UseBasicParsing | ConvertFrom-Json)[0].tag_name

		$c = secman verx

		if ($l -ne $c) {
			$nr = "there's a new release of secman is avalaible:"
			$up = "to upgrade run "
			$smu = "sm-upg start"

			Write-Host ""
			Write-Host -NoNewline $nr -ForegroundColor DarkYellow
			Write-Host "$c -> $l" -ForegroundColor DarkCyan
			Write-Host -NoNewline $up -ForegroundColor DarkYellow
			Write-Host $smu -ForegroundColor DarkCyan
		}
	`
}
	
func Check_ml() string {
	return `
		l=$(curl --silent "https://api.github.com/repos/secman-team/secman/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
		c=$(secman verx | tr -d '\n')

		if [ $l != $c ]; then
			nr="there's a new release of secman is avalaible:"
			up="to upgrade run"
			smu="secman upgrade"

			echo ""
			echo "$nr $c -> $l"
			echo "$up $smu"
		fi
	`
}
