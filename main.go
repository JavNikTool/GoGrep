package main

import (
	"flag"
	"log"
	"os"

	gflag "github.com/JavNikTool/GoGrep/flag"
	"github.com/JavNikTool/GoGrep/search"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatal("Должно быть 2 аргумента")
	}
	// структура для хранения состояния флагов
	flagList := gflag.FlagList{}
	// набор флагов
	flag.BoolVar(&flagList.IgnoreCase, "i", false, "игнорировать регистр")

	flag.Parse()
	// аргументы из cli
	path := flag.Arg(0)
	needle := flag.Arg(1)

	file, err := os.Stat(path)

	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	if !file.IsDir() {
		search.InFile(path, needle, flagList)
	} else {
		search.InDirectory(path, needle, flagList)
	}

}
