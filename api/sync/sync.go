package sync

import (
	"fmt"
	"os/exec"
)

func PushSync() {
	fmt.Println("syncing...")
	cmd := exec.Command("secman-sync", "phx")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
