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
		checkError(err)
		fmt.Printf("%v>", dir)
		waitUser.Scan()
		command := strings.Split(waitUser.Text(), " ")
		if len(command)-1 == 0 {
			choice(command[0], "")
		} else {
			choice(command[0], command[1])
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
			checkError(err)
		} else {
			err := LsFunc(".")
			checkError(err)
		}

	case "clear":
		for i := 0; i <= 200; i++ {
			fmt.Println()
		}
	case "show":
		ShowOpen(arg)
	case "cd":
		Cd(arg)
	default:
		fmt.Println("Неизвестная команда")
	}
}

// Т.к проверок на ошибки будет множество лучше вывести проверку ошибок в отдельную функцию?
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
