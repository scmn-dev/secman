package upg

import "github.com/secman-team/shell"

func Upgrade() {
	shell.ShellCmd("verx --upg")
}
