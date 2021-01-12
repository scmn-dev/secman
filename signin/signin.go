package signin

import (
	"os/exec"
	"fmt"
)

func Auth() {
  cmd := exec.Command("gh", "auth", "login")
  stdout, err := cmd.Output()

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Print(string(stdout))
}
