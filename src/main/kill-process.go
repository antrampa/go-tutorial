package main

import (
	"os/exec"
	"syscall"
	"bytes"
	"fmt"
)

func main() {
	isRunning, err := isProcRunning("java.exe")
	fmt.Println("errors:")
	fmt.Println(err)
	fmt.Println("is running java.exe: ")
	fmt.Println(isRunning)

	if isRunning {
		//Kill this task
		// Taskkill /IM firefox.exe /F
		killed, errK := KillTaskByName("java.exe")

		if errK != nil{
			fmt.Println(errK)
		}

		if killed {
			fmt.Println("java.exe killed ")
		}
	}
}

func isProcRunning(names ...string) (bool, error) {
	if len(names) == 0 {
		return false, nil
	}

	cmd := exec.Command("tasklist.exe", "/fo", "csv", "/nh")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}

	for _, name := range names {
		if bytes.Contains(out, []byte(name)) {
			return true, nil
		}
	}
	return false, nil
}

func KillTaskByName(name string) (bool, error) {
	if len(name) == 0 {
		return false, nil
	}

	cmd := exec.Command("Taskkill", "/IM", name, "/F")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}

	if out != nil {
		return true, nil
	}

	return false, nil
}