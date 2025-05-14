package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
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
	if s == "" || s == "[]" {
		return []int{}
	}
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
	if s == "" || s == "{}" {
		return map[string][]int{}
	}
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
	if s == "" || s == "{}" {
		return map[int]int{}
	}
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
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("ошибка конвертации строки %q в число int\n", s)
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
		t = append(t, test)
	}

	if err := sc.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла")
		os.Exit(1)
	}
	return t
}

var tests []testsType

func init() {
	tests = fillTests("numbers.csv")
}

func TestFreq(t *testing.T) { // TestFreq
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getFreq(tt.input)              // getFreq
			if !reflect.DeepEqual(tt.freq, ans) { // tt.freq
				t.Errorf("got %v, expected %v", ans, tt.freq) // tt.freq
			}
		})
	}
}

func TestMin(t *testing.T) { // TestFreq
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getMin(tt.input)              // getFreq
			if !reflect.DeepEqual(tt.min, ans) { // tt.freq
				t.Errorf("got %v, expected %v", ans, tt.min) // tt.freq
			}
		})
	}
}

func TestMax(t *testing.T) { // TestFreq
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getMax(tt.input)              // getFreq
			if !reflect.DeepEqual(tt.max, ans) { // tt.freq
				t.Errorf("got %v, expected %v", ans, tt.max) // tt.freq
			}
		})
	}
}

func TestAvg(t *testing.T) { // TestFreq
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getAvg(tt.input)              // getFreq
			if !reflect.DeepEqual(tt.avg, ans) { // tt.freq
				t.Errorf("got %v, expected %v", ans, tt.avg) // tt.freq
			}
		})
	}
}

func TestMode(t *testing.T) { // TestFreq
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getMode(tt.input)              // getFreq
			if !reflect.DeepEqual(tt.mode, ans) { // tt.freq
				t.Errorf("got %v, expected %v", ans, tt.mode) // tt.freq
			}
		})
	}
}

func TestEven(t *testing.T) { // TestFreq
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getEven(tt.input)              // getFreq
			if !reflect.DeepEqual(tt.even, ans) { // tt.freq
				t.Errorf("got %v, expected %v", ans, tt.even) // tt.freq
			}
		})
	}
}

func TestOdd(t *testing.T) { // TestFreq
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getOdd(tt.input)              // getFreq
			if !reflect.DeepEqual(tt.odd, ans) { // tt.freq
				t.Errorf("got %v, expected %v", ans, tt.odd) // tt.freq
			}
		})
	}
}

func TestRanges(t *testing.T) { // TestFreq
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getRanges(tt.input)              // getFreq
			if !reflect.DeepEqual(tt.ranges, ans) { // tt.freq
				t.Errorf("got %v, expected %v", ans, tt.ranges) // tt.freq
			}
		})
	}
}

func TestMedian(t *testing.T) { // TestFreq
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getMedian(tt.input)              // getFreq
			if !reflect.DeepEqual(tt.median, ans) { // tt.freq
				t.Errorf("got %v, expected %v", ans, tt.median) // tt.freq
			}
		})
	}
}

func TestUnique(t *testing.T) { // TestFreq
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getUnique(tt.input)              // getFreq
			if !reflect.DeepEqual(tt.unique, ans) { // tt.freq
				t.Errorf("got %v, expected %v", ans, tt.unique) // tt.freq
			}
		})
	}
}

func TestAboveAvg(t *testing.T) { // TestFreq
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getAboveAvg(tt.input)              // getFreq
			if !reflect.DeepEqual(tt.aboveAvg, ans) { // tt.freq
				t.Errorf("got %v, expected %v", ans, tt.aboveAvg) // tt.freq
			}
		})
	}
}
