package main

import (
	// "testing"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type testsType struct {
	name     string
	input    []int
	size     int
	freq     map[int]int
	min      int
	max      int
	avg      int
	mode     []int
	even     []int //четные
	odd      []int //нечетные
	ranges   map[string][]int
	median   int
	unique   []int
	aboveAvg []int
}

func parseSlice(s string) []int {
	sClean := strings.Replace(s, "[", "", -1)
	sClean = strings.Replace(sClean, "]", "", -1)
	slice := strings.Split(sClean, ";")
	rez := []int{}
	var val int
	for _, v := range slice {
		val = convertInt(v)
		rez = append(rez, val)
	}
	return rez
}

func parseDictStrSlInt(s string) map[string][]int {
	sClean := strings.Replace(s, "{", "", -1)
	sClean = strings.Replace(sClean, "}", "", -1)
	slice := strings.Split(sClean, "|")
	rez := map[string][]int{}
	var p []string
	var key string
	var val []int
	for _, v := range slice {
		p = strings.Split(v, ":")
		key = p[0]
		val = parseSlice(p[1])
		rez[key] = val
	}
	return rez
}

func parseDictIntInt(s string) map[int]int {
	sClean := strings.Replace(s, "{", "", -1)
	sClean = strings.Replace(sClean, "}", "", -1)
	slice := strings.Split(sClean, ";")
	rez := map[int]int{}
	var p []string
	var key, val int

	for _, v := range slice {
		p = strings.Split(v, ":")
		key = convertInt(p[0])
		val = convertInt(p[1])
		rez[key] = val
	}
	return rez
}

func convertInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("ошибка конвертации")
		os.Exit(1)
	}
	return i
}

func fillTests(path string) []testsType {

	t := []testsType{}

	f, err := os.Open(path)

	if err != nil {
		fmt.Printf("Ошибка при открытии файла")
		os.Exit(1)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var line string
	test := testsType{}
	var lineSlice []string
	i := 0
	for sc.Scan() {
		if i == 0 {
			i++
			continue
		}
		line = sc.Text()
		lineSlice = strings.Split(line, ",")
		test = testsType{
			name:     lineSlice[0],
			input:    parseSlice(lineSlice[1]),
			size:     convertInt(lineSlice[2]),
			freq:     parseDictIntInt(lineSlice[3]),
			min:      convertInt(lineSlice[4]),
			max:      convertInt(lineSlice[5]),
			avg:      convertInt(lineSlice[6]),
			mode:     parseSlice(lineSlice[7]),
			even:     parseSlice(lineSlice[8]), //четные
			odd:      parseSlice(lineSlice[9]), //нечетные
			ranges:   parseDictStrSlInt(lineSlice[10]),
			median:   convertInt(lineSlice[11]),
			unique:   parseSlice(lineSlice[12]),
			aboveAvg: parseSlice(lineSlice[13]),
		}
		fmt.Printf("%#v\n", test)
		t = append(t, test)
	}

	if err := sc.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла")
		os.Exit(1)
	}
	return t
}

func main() {
	tests := fillTests("numbers.csv")
	fmt.Printf("%v", tests)
}
