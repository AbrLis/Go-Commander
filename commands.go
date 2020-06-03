package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"path/filepath"
	"strings"
	"unicode/utf8"

	//"path/filepath"
	"os"
)

var Quit = false

func Check(err error) {
	if err != nil {
		fmt.Println("Ой! -> ", err)
		panic(err)
	}
}
func CheckErrFile(err error) {
	errColor := color.New(color.FgRed, color.Bold).Add(color.Underline)
	if pe, ok := err.(*os.PathError); ok {
		errColor.Printf("Ошибка: %s!\n", pe.Err)
	}
}

// Вывести список файлов и папок на экран
func LsFunc(path string) error {
	path, err := filepath.Abs(path)
	Check(err)
	files, err := ioutil.ReadDir(path)
	CheckErrFile(err)
	fileColor := color.New(color.FgYellow)                              // устанавливаем желтый цвет для файлов
	dirColor := color.New(color.FgGreen, color.Bold).Add(color.BgBlack) // устанавливаем выделение и зеленый цвет для папок

	//Находим самую длинную строку файла
	var lenFile int
	for _, file := range files {
		if len(file.Name()) > lenFile {
			lenFile = utf8.RuneCountInString(file.Name())
		}
	}
	showFile(dirColor, files, lenFile, true)   //Печать директорий
	showFile(fileColor, files, lenFile, false) //Печать файлов
	fmt.Println()                              // после вывода списка файлов и папок переводим курсор на новую строку
	return nil
}

//Вывод на консоль списка файлов
func showFile(fileColor *color.Color, files []os.FileInfo, lenFile int, flag bool) {
	nn := 80 / (lenFile + 2) //Допустимая длинна строки имени файла из расчёта 80 символов
	count := nn
	for _, file := range files {
		formatFile := fmt.Sprintf("/%s", file.Name()) //Формат вывода для директорий
		if !flag {
			formatFile = fmt.Sprintf("%s", file.Name()) //Формат вывода для файлов
		}
		if file.IsDir() == flag {
			_, err := fileColor.Printf(formatFile)
			Check(err)
			fmt.Print(strings.Repeat(" ", lenFile-utf8.RuneCountInString(file.Name())+2))
			count--
			if count == 0 {
				count = nn
				fmt.Println()
			}
		}
	}
}

// Вывести содержимое файла на консоль
func ShowOpen(file string) {
	file, err := filepath.Abs(file)
	Check(err)
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

// Сменить директорию
func Cd(path string) {
	path, err := filepath.Abs(path)
	Check(err)
	err = os.Chdir(path)
	CheckErrFile(err)
}

// Создать директорию
func MakeDir(name string) {
	if err := os.Mkdir(name, 0600); err != nil {
		fmt.Println("Имя директории содержит ошибки")
	}
}

func DeleteDir(path string) {
	path, err := filepath.Abs(path)
	Check(err)
	err = os.RemoveAll(path)
	if err != nil {
		fmt.Println("Имя директории содержит ошибки", err)
	}
}

// Переименовать файл или директорию
func Rename(name string) {
	file := strings.Split(name, " ")
	if len(file) < 2 {
		fmt.Println("Нет второго аргумента!")
		return
	}
	if file[1] == "" {
		fmt.Println("Второй аргумент не может быть пустым")
		return
	}
	err := os.Rename(file[0], file[1])
	Check(err)
}
