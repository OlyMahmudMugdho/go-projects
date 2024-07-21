package main

import (
	"fmt"
	"log"
	"os"
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

func GetProcess(pid int) (*os.Process, error) {
	process, err := os.FindProcess(pid)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return process, nil
}
