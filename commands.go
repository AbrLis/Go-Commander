package main
import (
	"fmt"
	//"path/filepath"
	"os"
	"bufio"
	//"strings"
	"io/ioutil"
	"github.com/fatih/color"
)

var Quit = false

func LsFunc(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		
		d := color.New(color.FgYellow, color.Bold)
		d.Printf("%s\t", file.Name())
	}
	fmt.Println()
	
	return nil
}

func ShowOpen(file string) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
	}
	defer readFile.Close()
	scanFile := bufio.NewScanner(readFile)
	for scanFile.Scan() {
		fmt.Println(scanFile.Text())
	}
}

func Cd(path string) {
	if err := os.Chdir(path); err != nil {
		fmt.Println("Ошибка чтения директории")
	}
}