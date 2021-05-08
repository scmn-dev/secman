package upg

import (
	"fmt"
	"runtime"

	"github.com/secman-team/shell"
	commands "github.com/secman-team/secman/tools/constants"
)

func Upgrade() {
	if runtime.GOOS == "windows" {
		fmt.Println("run sm-upg start")

	} else {
		shell.ShellCmd(commands.Upgrade())
	}
}
