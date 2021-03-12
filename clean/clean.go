package clean

import (
	"runtime"
	"github.com/secman-team/shell"
)

func Clean() {
	if runtime.GOOS == "windows" {
		shell.PWSLCmd("~/sm/clean.ps1")
	} else {
		shell.ShellCmd("ruby /home/sm/clean.rb")
	}
}
