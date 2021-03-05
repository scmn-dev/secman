package checker

import (
	"time"
	"github.com/briandowns/spinner"
	"github.com/secman-team/shell"
)

func Checker() {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Checking for updates..."
	s.Start()

	shell.SHCore("verx --sm", "bash $HOME/sm/vx --sm")

	s.Stop()
}
