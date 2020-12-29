package main

import "os/exec"
import . "fmt"

func main() {
  gh := "gh"

  // arg0 := "-e"
  // arg1 := "Hello world"
  // arg2 := "\n\tfrom"
  // arg3 := "golang"

  cmd := exec.Command(gh)
  stdout, err := cmd.Output()

  if err != nil {
    Println(err.Error())
    return
  }

  Print(string(stdout))
}
