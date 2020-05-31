package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	//"path/filepath"
	"os"
)

var Quit = false

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func LsFunc(path string) error {
	files, err := ioutil.ReadDir(path)
	Check(err)

	fileColor := color.New(color.FgYellow, color.Bold)
	dirColor := color.New(color.FgGreen, color.Bold).Add(color.BgBlack)

	for _, file := range files {
		if !file.IsDir() {
			_, err := fileColor.Printf("%s", file.Name())
			Check(err)
		} else {
			_, err := dirColor.Printf("[%s]", file.Name())
			Check(err)
		}
		fmt.Print("\t")
	}

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
