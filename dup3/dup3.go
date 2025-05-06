package main

import (
	"fmt"
	"os"
	"strings"
	//"bufio"
)

func main() {
	filesNames := os.Args[1:]
	counts := make(map[string]int)
	for _, fileName := range filesNames {
		fileContent, err := os.ReadFile(fileName)
		//fmt.Println(fileContent)
		if err == nil {

			fileContentSlice := strings.Split(string(fileContent), "\n")
			for line, lineContent := range fileContentSlice {
				if line == len(fileContentSlice)-1 {
					break
				}
				//fmt.Printf("%q\n", lineContent)
				counts[lineContent]++
			}
		}

	}
	//fmt.Println(counts)
	for line, countsValue := range counts {
		if countsValue > 1 {
			fmt.Printf("%d\t%s\n", countsValue, line)
		}
	}
}
