package upg

import (
	"fmt"
	"runtime"

	"github.com/secman-team/shell"
)

upgrade := `
	l=$(curl --silent "https://api.github.com/repos/secman-team/secman/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
	c=$(secman verx)
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

func Upgrade() {
	if runtime.GOOS == "windows" {
		// shell.PWSLCmd("& ~/sm/vx.ps1 --upg")
		fmt.Println("run sm-upg start")
	} else {
		shell.ShellCmd(upgrade)
	}
}
