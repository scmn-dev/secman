package fetch

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/secman-team/shell"
)

func FetchSECDIR() {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Suffix = " Fetching..."
	s.Start()

	shell.SHCore("secman-sync pl", "bash $HOME/sm/ssc pl")

	s.Stop()
}
