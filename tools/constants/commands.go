package commands

func Fetch_w() string {
	return `
		Remove-Item $HOME/.secman -Recurse -Force
		secman sync fetch-clone

		Write-Host "Fetched Successfully"
	`
}

func Fetch_ml() string {
	return `
		cd $HOME/.secman
		git pull
		echo "Fetched Successfully ✅"
	`
}

func Upgrade() string {
	return `
		l=$(curl --silent "https://api.github.com/repos/scmn-dev/secman/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
		c=$(secman verx | tr -d \n)
		smLoc="/usr/local/bin/secman*"

		if [ $l == $c ]; then
			echo "secman is already up-to-date and it's the latest release $l"

		elif [ $l != $c ]; then
			sudo rm $smLoc

			curl -fsSL https://unix-upg.secman.dev | bash

			if [ -x "command -v $(secman)" ]; then
				echo "secman was upgraded successfully 🎊"
			fi
		fi
	`
}

func Uninstall() string {
	return `
		smLoc=/usr/local/bin/secman*
		smManLoc=/usr/share/man/man1/secman*.1.gz

		rmMain() {
			if [ -x "$(command -v sudo)" ]; then
				sudo rm $smManLoc
				sudo rm $smLoc
			else
				rm $smManLoc
				rm $smLoc
			fi

			rm -rf ~/.sm-cluster
		}

		if [ -x "$(command -v secman)" ]; then
			rmMain

			if ! [ -f "$smLoc" ]; then
				echo "secman was uninstalled successfully... thank you for using secman 👋"

			else
				echo "there's an error while uninstalling secman, try again"
			fi

		else
			echo "there's no secman 😐"
		fi
	`
}

func ClearData() string {
	return `
		SECDIR=~/.secman

		afterClear() {
			username=$(secman auth get-username)

			if ! [ -d $SECDIR ]; then
				echo "after clear, if you want to restore .secman you can clone it from your private repo in https://github.com/$username/.secman"
			fi
		}

		if [ -x "$(command -v sudo)" ]; then
			sudo rm -rf $SECDIR

			afterClear
		else
			rm -rf $SECDIR

			afterClear
		fi
	`
}

func Clean_w() string {
	return `
		$directoyPath = "$HOME\.secman"

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
		#!/bin/bash
		if [ -d $HOME/.secman ]; then rm -rf $HOME/.secman; fi
		if ! [ -d $HOME/.secman ]; then echo "secman was cleaned successfully 🧹"; fi
	`
}

func Clean_w_git() string {
	return `
		$directoyPath = "$HOME\.secman"

		if (Test-Path -path $directoyPath) {
			Remove-Item $directoyPath/.git -Recurse -Force
		}

		if (!(Test-Path -path $directoyPath/.git)) {
			Write-Host "git was removed"
		}
	`
}

func Clean_ml_git() string {
	return `
		#!/bin/bash
		if [ -d $HOME/.secman/.git ]; then rm -rf $HOME/.secman/.git; fi
		if ! [ -d $HOME/.secman/.git ]; then echo "git was removed"; fi
	`
}

func Start_w() string {
	return `
		$username = secman auth get-username
		$SECDIR = $HOME\.secman
		cd $SECDIR

		git init

		Write-Host "# My secman passwords - $username" >> $SECDIR\README.md

		secman repo create .secman -d "My secman passwords - $username" --private -y

		git add .
		git commit -m "new .secman repo"
		git branch -M trunk
		git remote add origin https://github.com/$username/.secman
		git push -u origin trunk

		cd $lastDir
	`
}

func Start_ml() string {
	return `
		username=$(secman auth get-username)
		cd $HOME/.secman
		git init

		echo "# My secman passwords - $username" >> $HOME/.secman/README.md

		secman repo create .secman -d "My secman passwords - $username" --private -y

		git add .
		git commit -m "new .secman repo"
		git branch -M trunk
		git remote add origin https://github.com/$username/.secman
		git push -u origin trunk
	`
}

func StartEX() string {
	return "echo '\n## Clone\n\n```\nsecman sync clone\n```\n\n## Open Your Repo\n\n```\nsecman open\n```\n\n> Open your repo in the browser\n\n```\nsecman open -w/--web\n```\n\n**for more about sync command, run `secman sync -h`**\n**or visit the docs at https://secman.dev/docs/commands/sync**' >> $HOME/.secman/README.md"
}

func Push_w() string {
	return `
		$lastDir = pwd
		cd $HOME\.secman

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
		cd $HOME/.secman
		git add .
		git commit -m "new secman password"
		git push
	`
}

func Pull_w() string {
	return `
		$lastDir = pwd
		cd $HOME\.secman
		
		git pull
		
		cd $lastDir
	`
}

func Pull_ml() string {
	return `
		cd $HOME/.secman
		git pull
	`
}

func Clone() string {
	return `
		secman repo clone .secman $HOME/.secman
	`
}

func Clone_check_w() string {
	return `
		if (Test-Path -path $HOME/.secman) {
			Write-Host "cloned successfully"
		}
	`
}

func Clone_check_ml() string {
	return `if [ -d $HOME/.secman ]; then echo "cloned successfully ✅"; fi`
}

func Clone_Template() string {
	return "git clone https://github.com/scmn-dev/.secman.ex $HOME/.secman"
}

func Check_w() string {
	return `
		$releases = "https://api.github.com/repos/scmn-dev/secman/releases"

		$l = (Invoke-WebRequest -Uri $releases -UseBasicParsing | ConvertFrom-Json)[0].tag_name

		$c = secman verx

		if ($l -ne $c) {
			$nr = "there's a new release of secman is avalaible: "
			$up = "to upgrade run "
			$smu = "sm-win start"

			Write-Host ""
			Write-Host -NoNewline $nr -ForegroundColor DarkYellow
			Write-Host "$c → $l" -ForegroundColor DarkCyan
			Write-Host -NoNewline $up -ForegroundColor DarkYellow
			Write-Host $smu -ForegroundColor DarkCyan
		}
	`
}

func Check_ml() string {
	return `
		l=$(curl --silent "https://api.github.com/repos/scmn-dev/secman/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
		c=$(secman verx | tr -d \n)

		if [ $l != $c ]; then
			nr="there's a new release of secman is avalaible: "
			up="to upgrade run"
			smu="sudo secman upgrade"

			echo ""
			echo "$nr $c → $l"
			echo "$up $smu"
		fi
	`
}
