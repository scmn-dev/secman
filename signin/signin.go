package signin

import "os/exec"
import . "fmt"

func main() {
  gh := "gh"
  auth := "auth"
  login := "login"

  cmd := exec.Command(gh, auth, login)
  stdout, err := cmd.Output()

  if err != nil {
    Println(err.Error())
    return
  }

  Print(string(stdout))
}
