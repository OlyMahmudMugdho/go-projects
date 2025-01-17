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
	RenameFile("some.txt", "random.txt")
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

func KillProcess(process *os.Process) error {
	return process.Kill()
}

func RenameFile(oldName string, newName string) error {
	return os.Rename(oldName, newName)
}
