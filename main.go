package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func searchNeedleInFile(path, needle string) {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal("Ошибка при чтении файла")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for lineNum := 1; scanner.Scan(); lineNum++ {
		line := scanner.Text()

		if strings.Contains(line, needle) {
			fmt.Printf("%v:%v:%v\n", path, lineNum, line)
		}
	}

}

func main() {

	if len(os.Args) < 3 {
		log.Fatal("Должно быть 2 аргумента")
	}

	path := os.Args[1]
	needle := os.Args[2]

	file, err := os.Stat(path)

	if err != nil {
		log.Fatal("Файл или папка не найдены")
	}

	if !file.IsDir() {
		searchNeedleInFile(path, needle)
	}

}
