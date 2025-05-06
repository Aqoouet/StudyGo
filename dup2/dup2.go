package main

import (
	"fmt"
	"bufio"
	"os"
	"slices"
)

func main() {
	filesNames := os.Args[1:]
	counts:= make(map[string]int)
	lineFileNames := make(map[string][]string)
	if len(filesNames) == 0 {
		countLines(os.Stdin,counts,lineFileNames, "введено вручную")
	} else {
		for _, fileName := range filesNames {
			file, err := os.Open(fileName)
			if err!=nil {
				fmt.Fprintf (os.Stderr, "%v\n", err)
			} else {
				countLines(file,counts,lineFileNames, fileName)
			}
			defer file.Close()
		}
		for line,countValue := range counts{
			if countValue > 1 {
				fmt.Printf("%d\t%s\t%s\n", countValue, line, lineFileNames[line])
			}
		}
	}
}

func countLines(streamF *os.File, counts map[string]int, lineFileNames map[string][]string, fileName string) {
	input := bufio.NewScanner(streamF)
	for input.Scan() {
		currentLine:=input.Text()
		counts[currentLine]++
		if !slices.Contains(lineFileNames[currentLine], fileName) {
			lineFileNames[currentLine] = append(lineFileNames[currentLine], fileName)
		}
	}
}
