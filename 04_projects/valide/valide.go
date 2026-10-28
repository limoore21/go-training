package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//type File struct {
//	id    int
//	name  string
//	group string
//}

func main() {
	file, err := os.Open("valide.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	resFile, err := os.Create("rezult.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resFile.Close()
	writer := bufio.NewWriter(resFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Разработчики") {
			fmt.Println("Найдена строка с разработчиками:", line)
			writer.WriteString(line + "\n")
		}
	}
}
