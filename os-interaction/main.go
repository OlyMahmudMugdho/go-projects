package main

import (
	"fmt"
	"os/exec"
)

func main() {

	username, _ := RunCommand("whoami")
	fmt.Println(string(username))
}

func RunCommand(command string) ([]byte, error) {
	foundCommand := exec.Command(command)
	return foundCommand.Output()
}
