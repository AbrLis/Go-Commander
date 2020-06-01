package main

import (
	"fmt"
	//"path/filepath"
	"bufio"
	"os"
	"strings"
	//"io/ioutil"
)

func main() {
	waitUser := bufio.NewScanner(os.Stdin)
	for {
		dir, err := os.Getwd()
		Check(err)
		fmt.Printf("%v>", dir)
		waitUser.Scan()
		command := strings.Split(waitUser.Text(), " ")
		if len(command)-1 == 0 {
			choice(command[0], "")
		} else {
			choice(command[0], CutFirstString(command[0], waitUser.Text()))
		}
		if Quit {
			return
		}
	}
}

func choice(com string, arg string) {
	switch com {
	case "exit":
		Quit = true
		return
	case "ls":
		if arg != "" {
			err := LsFunc(arg)
			Check(err)
		} else {
			err := LsFunc(".")
			Check(err)
		}

	case "clear":
		for i := 0; i <= 200; i++ {
			fmt.Println()
		}
	case "show":
		ShowOpen(arg)
	case "cd":
		Cd(arg)
	case "mkdir":
		MakeDir(arg)
	case "rmdir":
		DeleteDir(arg)
	default:
		fmt.Println("Неизвестная команда")
	}
}
