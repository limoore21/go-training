package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	pro      sync.WaitGroup
	penis    sync.Mutex
	count    = make(map[string]int)
	pifFiles = make(map[string][]string)
	dsadas   []os.DirEntry
)

func main() {
	dirpath := "/home/max/Загрузки"
	var err error
	dsadas, err = os.ReadDir(dirpath)
	if err != nil {
		fmt.Println(err)
	}
	for _, d := range dsadas {
		if d.IsDir() {
			continue
		}
		pro.Add(1)
		go Filepro(d)
	}
	pro.Wait()
	fmt.Println("===Статистика по расширениям===")
	for pif, files := range pifFiles {
		printExt := pif
		if printExt == "" {
			printExt = "[Без расширения]"
		}
		fmt.Printf("%s: %d файл(ов)\n", printExt, len(files))
		fmt.Printf("   Файлы: %s\n", strings.Join(files, ", "))
	}
}
func Filepro(d os.DirEntry) {
	defer pro.Done()
	ext := filepath.Ext(d.Name())
	penis.Lock()
	count[ext]++
	pifFiles[ext] = append(pifFiles[ext], d.Name())
	penis.Unlock()
}
