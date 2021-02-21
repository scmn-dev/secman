package fetch

import (
	"fmt"
	"os/exec"
	"github.com/briandowns/spinner"
	"time"
)

func Checker() {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Suffix = " Fetching..."
	s.Start()
	
	cmd := exec.Command("secman-sync", "pl")
	stdout, err := cmd.Output()
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	fmt.Print(string(stdout))
	s.Stop()
}
