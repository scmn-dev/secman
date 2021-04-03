package login

import (
	"fmt"
	"runtime"

	"github.com/secman-team/shell"
)

func LoginCore() {
	const cmd string = "gh auth login"

	if runtime.GOOS == "windows" {
		shell.PWSLCmd(cmd)
	} else {
		shell.ShellCmd(cmd)
	}
}
