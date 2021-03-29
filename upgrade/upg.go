package upg

import (
	"github.com/secman-team/shell"
	"runtime"
)

func Upgrade() {
	if runtime.GOOS == "windows" {
		shell.PWSLCmd("& ~/sm/vx.ps1 --upg")
	} else {
		shell.ShellCmd("verx --upg")
	}
}
