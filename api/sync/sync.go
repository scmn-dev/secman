package sync

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"github.com/secman-team/shell"
)

func PushSync() {
	const Syncing string = " 📮 Syncing..."

	if runtime.GOOS == "windows" {
		err, out, errout := shell.PWSLOut("& $HOME/sm/sync.ps1")

		fmt.Print(out)

		if err != nil {
			log.Printf("error: %v\n", err)
			fmt.Print(errout)
		} else if out != "" {
			s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
			s.Suffix = Syncing
			s.Start()

			shell.PWSLCmd("& $HOME/sm/secman-sync.ps1 ph")

			s.Stop()
		}
	} else {
		err, out, errout := shell.ShellOut("/home/sm/sync.sh")

		fmt.Print(out)

		if err != nil {
			log.Printf("error: %v\n", err)
			fmt.Print(errout)
		} else if out != "" {
			s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
			s.Suffix = Syncing
			s.Start()

			shell.ShellCmd("secman-sync phx")

			s.Stop()
		}
	}
}
