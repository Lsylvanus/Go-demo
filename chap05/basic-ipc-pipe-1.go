package main

import (
	"os/exec"
	"fmt"
)

func main() {
	//create cmd
	cmd_go_env := exec.Command("go", "env")
	//cmd_grep := exec.Command("grep","GOROOT")

	stdout_env, env_error := cmd_go_env.StdoutPipe()
	if env_error != nil {
		fmt.Println("Error happened about standard output pipe ", env_error)
		return
	}

	//env_error := cmd_go_env.Start()
	if env_error := cmd_go_env.Start(); env_error != nil {
		fmt.Println("Error happened in execution ", env_error)
		return
	}

	a1 := make([]byte, 1024)
	n, err := stdout_env.Read(a1)
	if err != nil {
		fmt.Println("Error happened in reading from stdout", err)
	}

	fmt.Printf("Standard output of go env command: %s", a1[:n])
}