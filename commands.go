package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"strings"
	"unicode/utf8"

	//"path/filepath"
	"os"
)

type CmdData struct {
	command    string // Команда
	firstFile  string //
	SecondFile string //
	firstPath  string //
	secondPath string //
	trash      string // Остаток нераспознанной строки
}

// Парсинг введённой строки и разбивка по данным в CmdData
// Скорее всего множество ошибок и недочётов
func (c *CmdData) ParseCommand(str string) {
	//Обработка команды
	str = strings.TrimPrefix(str, " ")
	idx := strings.IndexAny(str, " ")
	if idx == -1 {
		c.command = strings.ToLower(str)
		c.firstPath = "."
		return
	} else {
		c.command = strings.ToLower(str[:idx])
		str = str[idx:]
	}
	str = strings.TrimPrefix(str, " ")
	//Обработка путей и файлов
	c.firstPath, str = GetMePath(str)
	c.firstFile, str = GetMeFileName(str)
	if c.firstFile == "" && c.firstPath == "." {
		c.trash = str
		return
	}
	//Обработка второго пути и файла
	c.secondPath, str = GetMePath(str)
	c.SecondFile, str = GetMeFileName(str)
	c.trash = str
}

//Ищет существующий путь в строке и возвращает его и остаток строки. Если не найдено то возвращает "."
func GetMePath(str string) (string, string) {
	tempDir, err := os.Getwd()
	Check(err)

	splitInput := strings.SplitAfter(str, " ")
	name := ""
	for _, v := range splitInput {
		name += v
		send := strings.TrimSuffix(name, " ")
		// Попытка прочитать директорию
		if err := os.Chdir(send); err != nil {
			continue
		} else {
			//Директория существует
			err := os.Chdir(tempDir)
			Check(err)
			return send, CutFirstString(send, strings.TrimPrefix(str, " "))
		}
	}
	// Не удалось найти существующую директорию
	err = os.Chdir(tempDir)
	Check(err)
	return ".", str
}

// Возвращает строку-файл если она присутствует в данной директории, остаток строки
func GetMeFileName(inputString string) (string, string) {
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
				return send, CutFirstString(send, inputString)
			}
		}
	}
	// Не удалось найти существующий файл
	return "", inputString
}

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

	if pe, ok := err.(*os.PathError); ok {
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
	file, _ = GetMeFileName(file)
	if file == "" {
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

// Сменить директорию
func Cd(path string) {
	if err := os.Chdir(path); err != nil {
		fmt.Println("Ошибка чтения директории")
	}
}

// Создать директорию
func MakeDir(name string) {
	if err := os.Mkdir(name, 0600); err != nil {
		fmt.Println("Имя директории содержит ошибки")
	}
}

func DeleteDir(prs CmdData) {
	del := ""
	if prs.firstFile != "" {
		del = prs.firstFile
	} else {
		del = prs.firstPath
	}
	err := os.RemoveAll(del)
	if err != nil {
		fmt.Println("Имя директории содержит ошибки", err)
	}
}

// Переименовать файл или директорию
func Rename(prs CmdData) {
	ren := ""
	if prs.firstFile != "" {
		ren = prs.firstFile
	} else {
		ren = prs.firstPath
	}
	err := os.Rename(ren, prs.trash)
	Check(err)
}

// Удаляет из строки 1 аргумент и подчищает впереди стоящие пробелы
// Добавил т.к встречается неоднократно и может понадобится в дальнейшем.
func CutFirstString(deleteString, originalString string) string {
	originalString = strings.Replace(originalString, deleteString, "", 1)
	originalString = strings.TrimPrefix(originalString, " ")
	return originalString
}
