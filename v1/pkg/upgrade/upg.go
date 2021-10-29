package upg

import (
	"fmt"
	"time"
	"runtime"

	"github.com/abdfnx/shell"
	commands "github.com/scmn-dev/secman-v1/tools/constants"
	"github.com/scmn-dev/secman-v1/tools/shared"

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
