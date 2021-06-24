package upg

import (
	"fmt"
	"time"
	"runtime"

	"github.com/secman-team/shell"
	commands "github.com/secman-team/secman/tools/constants"
	"github.com/secman-team/secman/tools/shared"

	"github.com/briandowns/spinner"
)

func Upgrade() {
	if runtime.GOOS == "windows" {
		fmt.Println(shared.RunSMWin())

	} else {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Suffix = " 🚧 Upgrading..."
		s.Start()

		shell.ShellCmd(commands.Upgrade())

		s.Stop()
	}
}
