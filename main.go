package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func searchNeedleInFile(path, needle string) {
	f, err := os.Open(path)

	if err != nil {
		log.Printf("Ошибка при чтении файла %s: %v", path, err)
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for lineNum := 1; scanner.Scan(); lineNum++ {
		line := scanner.Text()

		if strings.Contains(line, needle) {
			fmt.Printf("Файл: %s\n", path)
			fmt.Printf("%v:%v:%v\n", path, lineNum, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Ошибка при сканировании файла %s: %v", path, err)
	}
}

func searchNeedleInDirectory(path, needle string) {
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("предотвращение ошибки доступа к пути %s: %v\n", path, err)
			return nil
		}

		if d.IsDir() {
			fmt.Printf("Директория: %s\n", path)
			return nil
		} else {
			searchNeedleInFile(path, needle)
			return nil
		}
	})

	if err != nil {
		log.Fatalf("Ошибка при обходе директории: %v", err)
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
		log.Fatalf("Ошибка: %v", err)
	}

	if !file.IsDir() {
		searchNeedleInFile(path, needle)
	} else {
		searchNeedleInDirectory(path, needle)
	}

}
