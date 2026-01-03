package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FlagList struct {
	ignoreCase bool
}

func searchNeedleInFile(path, needle string, fl FlagList) {
	f, err := os.Open(path)

	if err != nil {
		log.Printf("Ошибка при чтении файла %s: %v", path, err)
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	if fl.ignoreCase {
		needle = strings.ToLower(needle)
	}

	for lineNum := 1; scanner.Scan(); lineNum++ {
		lineRaw := scanner.Text()

		lineCmp := lineRaw
		if fl.ignoreCase {
			lineCmp = strings.ToLower(lineRaw)
		}

		if strings.Contains(lineCmp, needle) {
			fmt.Printf("Файл: %s\n", path)
			fmt.Printf("%v:%v:%v\n", path, lineNum, lineRaw)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Ошибка при сканировании файла %s: %v", path, err)
	}
}

func searchNeedleInDirectory(path, needle string, fl FlagList) {
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("предотвращение ошибки доступа к пути %s: %v\n", path, err)
			return nil
		}

		if d.IsDir() {
			fmt.Printf("Директория: %s\n", path)
			return nil
		} else {
			searchNeedleInFile(path, needle, fl)
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
	// структура для хранения состояния флагов
	flagList := FlagList{}
	// набор флагов
	flag.BoolVar(&flagList.ignoreCase, "i", false, "игнорировать регистр")

	flag.Parse()
	// аргументы из cli
	path := flag.Arg(0)
	needle := flag.Arg(1)

	file, err := os.Stat(path)

	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	if !file.IsDir() {
		searchNeedleInFile(path, needle, flagList)
	} else {
		searchNeedleInDirectory(path, needle, flagList)
	}

}
