package main

import "os/exec"
import . "fmt"

func main() {
  cmd := exec.Command("gh", "repo")
  stdout, err := cmd.Output()

  if err != nil {
    Println(err.Error())
    return
  }

  Print(string(stdout))
}
