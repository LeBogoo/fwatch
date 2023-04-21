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
		var cmd *exec.Cmd

		if newModTime != modTime {
			fmt.Println("fwatch: File changed, running command...")
			modTime = newModTime

			if cmd != nil {
				cmd.Process.Kill()
				fmt.Println("fwatch: Killed previous command due to file change.")
			}

			cmd = exec.Command(os.Args[2], os.Args[3:]...)

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Start()

			if err != nil {
				errMessage := err.Error()
				// replace the "exec: " with nothing
				errMessage = errMessage[6:]
				fmt.Println("fwatch: Error: " + errMessage)
				os.Exit(1)
			}

		}

		time.Sleep(500 * time.Millisecond)
	}
}
