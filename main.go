package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func notfound() {
	fmt.Println("Error: File \"" + os.Args[1] + "\" not found.")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: fwatch <file> <command> <args>")
		return
	}

	filepath := os.Args[1]

	fileStat, err := os.Stat(filepath)

	if err != nil {
		notfound()
	}

	modTime := fileStat.ModTime()

	for {

		fileStat, err := os.Stat(filepath)

		if err != nil {
			notfound()
		}

		newModTime := fileStat.ModTime()

		if newModTime != modTime {
			modTime = newModTime
			cmd := exec.Command(os.Args[2], os.Args[3:]...)
			stdout, err := cmd.Output()

			if err != nil {
				errMessage := err.Error()
				// replace the "exec: " with nothing
				errMessage = errMessage[6:]
				fmt.Println("Error: " + errMessage)
				os.Exit(1)
			}

			fmt.Println(string(stdout))
		}

		time.Sleep(500 * time.Millisecond)
	}
}
