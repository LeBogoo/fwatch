package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	filepath := os.Args[1]
	fileStat, _ := os.Stat(filepath)

	modTime := fileStat.ModTime()

	for {

		fileStat, _ := os.Stat(filepath)
		newModTime := fileStat.ModTime()

		if newModTime != modTime {
			modTime = newModTime
			cmd := exec.Command(os.Args[2], os.Args[3:]...)
			stdout, err := cmd.Output()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println(string(stdout))
		}

		time.Sleep(500 * time.Millisecond)
	}
}
