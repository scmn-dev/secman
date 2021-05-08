package clean

import (
	"github.com/secman-team/shell"
	commands "github.com/secman-team/secman/tools/constants"
)

func Clean() {
	shell.SHCore(commands.Clean_ml(), commands.Clean_w())
}
