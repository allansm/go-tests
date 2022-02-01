package main

import "os/exec"

func main(){
	cmd := exec.Command("notepad")
	cmd.Run()
}
