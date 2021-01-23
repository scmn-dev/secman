package sync

import (
	"fmt"
	"os"
	"os/exec"
)

func PushSync() {
	if _, err := os.Stat("~/.secman/.git"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("syncing...")
			cmd := exec.Command("secman-sync", "phx")
			stdout, err := cmd.Output()
		
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		
			fmt.Print(string(stdout))
		}
	}
}
