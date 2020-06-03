package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	waitUser := bufio.NewScanner(os.Stdin)
	var command string
	for {
		dir, err := os.Getwd()
		Check(err)
		fmt.Printf("%v>", dir)
		waitUser.Scan()

		path := strings.TrimPrefix(waitUser.Text(), " ")
		idx := strings.IndexAny(path, " ")
		if idx == -1 {
			command = strings.ToLower(path)
			path = ""
		} else {
			command = strings.ToLower(path[:idx])
			path = path[idx:]
		}
		path = strings.TrimPrefix(path, " ")
		choice(command, path)

		//path = strings.TrimPrefix(path, " ")
		//var prs CmdData
		//prs.ParseCommand(waitUser.Text())
		//choice(prs)
		if Quit {
			return
		}
	}
}

func choice(command, path string) {
	switch command {
	case "exit":
		Quit = true
		return
	case "ls":
		err := LsFunc(path)
		Check(err)
	case "clear":
		for i := 0; i <= 200; i++ {
			fmt.Println()
		}
	case "show":
		ShowOpen(path)
	case "cd":
		Cd(path)
	case "mkdir":
		MakeDir(path)
	case "rmdir":
		DeleteDir(path)
	case "rename":
		Rename(path)
	default:
		fmt.Println("Неизвестная команда")
	}
}
