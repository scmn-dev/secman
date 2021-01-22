package sync

import (
	"fmt"
	"os/exec"
)

func PushSync() {
	cmd := exec.Command("secman-sync", "ph")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
