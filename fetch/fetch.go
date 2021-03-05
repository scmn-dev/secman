package fetch

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/secman-team/shell"
)

func FetchSECDIR() {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Suffix = " Fetching..."
	s.Start()

	SHCore("secman-sync pl", "bash ssc pl")

	s.Stop()
}
