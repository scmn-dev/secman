package upg

import "github.com/secman-team/shell"

func Upgrade() {
	shell.SHCore("verx --upg", "bash vx --upg")
}
