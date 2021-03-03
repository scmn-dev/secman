package checker

import (
	"fmt"
	"os/exec"
	"github.com/briandowns/spinner"
	"time"
)

func Checker() {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Checking for updates..."
	s.Start()
	
	cmd := exec.Command("verx", "--sm")
	stdout, err := cmd.Output()
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	fmt.Print(string(stdout))
	s.Stop()
}
