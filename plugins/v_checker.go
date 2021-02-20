package checker

import (
	"fmt"
	"os/exec"
)

func Checker() {
	cmd := exec.Command("bash", "~/sm/verx", "--sm")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("checking for updates...")
	fmt.Print(string(stdout))
}
