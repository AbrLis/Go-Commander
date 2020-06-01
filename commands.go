package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"strings"
	"unicode/utf8"

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
	errColor := color.New(color.FgRed, color.Bold).Add(color.Underline)

	if pe, ok := err.(*os.PathError);ok{
		errColor.Printf("Ошибка: %s!\n", pe.Err)
		//fmt.Printf("Op: %s!\n", pe.Op)
		//fmt.Printf("Path: %s\n", pe.Path)
	}

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
	file, _, err := GetMeFileName(file)
	if err != nil {
		fmt.Println("Файл не найден")
		return
	}
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

// Создать дирректорию
func MakeDir(name string) {
	if err := os.Mkdir(name, 0600); err != nil {
		fmt.Println("Имя дирректории содержит ошибки")
	}
}

// Возвращает строку директорию-файл если она присутствует в данной директории, остаток строки, код ошибки
func GetMeFileName(inputString string) (string, string, error) {
	inputString = strings.ToLower(inputString)
	files, err := ioutil.ReadDir(".")
	Check(err)

	splitInput := strings.SplitAfter(inputString, " ")
	name := ""
	for _, v := range splitInput {
		name += v
		send := strings.TrimSuffix(name, " ")
		for _, v := range files {
			if send == strings.ToLower(v.Name()) {
				return send, CutFirstString(send, inputString), nil
			}
		}
	}
	err = errors.New("file not found")
	return "", inputString, err
}

// Удаляет из строки 1 аргумент и подчищает впереди стоящие пробелы
// Добавил т.к встречается неоднократно и может понадобится в дальнейшем.
func CutFirstString(deleteString, originalString string) string {
	originalString = strings.Replace(originalString, deleteString, "", 1)
	originalString = strings.TrimPrefix(originalString, " ")
	return originalString
}

func DeleteDir(name string) {
	err := os.RemoveAll(name)
	if err != nil {
		fmt.Println("Имя дирректории содержит ошибки")
	}
}
