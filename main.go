package main

import (
	"fmt"
	//"path/filepath"
	"os"
	"bufio"
	"strings"
	//"io/ioutil"
)


func main() {
	for {
		waitUser := bufio.NewScanner(os.Stdin)
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
			if err := LsFunc(arg); err != nil {
				fmt.Printf("Какая-то ошибка: %v\n", err)
			}
		} else {
			if err := LsFunc("."); err != nil {
				fmt.Printf("Какая-то ошибка: %v\n", err)
			}
		}
		
	case "clear":
		for i:=0;i<=200;i++ {
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