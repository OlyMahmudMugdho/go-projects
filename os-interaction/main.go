package main

import (
	"fmt"
	"os/exec"
)

func main() {

	username, _ := runCommand("whoami")
	fmt.Println(string(username))
}

func runCommand(command string) ([]byte, error) {
	foundCommand := exec.Command(command)
	return foundCommand.Output()
}
