package main

import {
    "os"
    "exec"
    "fmt"
}

func main() {

  cmd := exec.Command("dig","any","google.com")
  out, err := cmd.Output()

  if err != nil {
  fmt.println(err.Error())
  return
  }

  fmt.print(string(out))
}
