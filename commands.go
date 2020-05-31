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

// Вывести список файлов и папок на экран
func LsFunc(path string) error {
	files, err := ioutil.ReadDir(path)
	Check(err)

	fileColor := color.New(color.FgYellow) // устанавливаем желтый цвет для файлов
	dirColor := color.New(color.FgGreen, color.Bold).Add(color.BgBlack) // устанавливаем выделение и зеленый цвет для папок

	for _, file := range files {
		if !file.IsDir() {
			_, err := fileColor.Printf("%s", file.Name()) // выводим список файлов
			Check(err)
		} else {
			_, err := dirColor.Printf("[%s]", file.Name()) // выводим список папок
			Check(err)
		}
		fmt.Print("\t") // после вывода списка файлов и папок переводим курсор на новую строку
	}
	fmt.Println()

	return nil
}

// Вывести содержимое файла на консоль
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

// Сменить дирректорию
func Cd(path string) {
	if err := os.Chdir(path); err != nil {
		fmt.Println("Ошибка чтения директории")
	}
}
