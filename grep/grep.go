package grep

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/JavNikTool/GoGrep/flag"
)

func InFile(path, needle string, fl flag.FlagList) {
	f, err := os.Open(path)

	if err != nil {
		log.Printf("Ошибка при чтении файла %s: %v", path, err)
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	if fl.IgnoreCase {
		needle = strings.ToLower(needle)
	}

	for lineNum := 1; scanner.Scan(); lineNum++ {
		lineRaw := scanner.Text()

		lineCmp := lineRaw
		if fl.IgnoreCase {
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

func InDirectory(path, needle string, fl flag.FlagList) {
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("предотвращение ошибки доступа к пути %s: %v\n", path, err)
			return nil
		}

		if d.IsDir() {
			fmt.Printf("Директория: %s\n", path)
			return nil
		} else {
			InFile(path, needle, fl)
			return nil
		}
	})

	if err != nil {
		log.Fatalf("Ошибка при обходе директории: %v", err)
	}
}
