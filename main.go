package main

import (
	"fmt"
	//"path/filepath"
	"bufio"
	"os"
	//"io/ioutil"
)

func main() {
	waitUser := bufio.NewScanner(os.Stdin)
	var prs CmdData
	for {
		dir, err := os.Getwd()
		Check(err)
		fmt.Printf("%v>", dir)
		waitUser.Scan()
		prs.ParseCommand(waitUser.Text())
		choice(prs)
		if Quit {
			return
		}
	}
}

func choice(prs CmdData) {
	switch prs.command {
	case "exit":
		Quit = true
		return
	case "ls":
		err := LsFunc(prs.firstPath)
		Check(err)
	case "clear":
		for i := 0; i <= 200; i++ {
			fmt.Println()
		}
	case "show":
		ShowOpen(prs.firstFile)
	case "cd":
		Cd(prs.firstPath)
	case "mkdir":
		MakeDir(prs.trash)
	case "rmdir":
		DeleteDir(prs.firstPath)
	default:
		fmt.Println("Неизвестная команда")
	}
}
