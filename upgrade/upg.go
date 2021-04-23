package upg

import (
	"fmt"
	"runtime"

	"github.com/secman-team/shell"
)

func Upgrade() {
	if runtime.GOOS == "windows" {
		// shell.PWSLCmd("& ~/sm/vx.ps1 --upg")
		fmt.Println("run sm-upg start")
	} else {
		shell.ShellCmd("verx --upg")
	}
}
