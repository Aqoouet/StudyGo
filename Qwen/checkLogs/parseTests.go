package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type TestCase struct {
	Name    string
	Descr   string
	Logs    []string
	Time1   time.Time
	Time2   time.Time
	Console []string
	CSV     []string
	JSON    string
}

func cutSlice(a []string) [][][]string {

	rez := [][][]string{}
	sub := [][]string{}

	for i := 0; i < len(a); {
		s := []string{}
		cat := strings.SplitN(a[i], ":", 2)[0]
		for i < len(a) && strings.Contains(a[i], cat) {
			s = append(s, strings.SplitN(a[i], ":", 2)[1])
			i++
		}
		sub = append(sub, s)
		if cat == "json" {
			rez = append(rez, sub)
			sub = [][]string{}
		}

	}
	return rez
}

func setCategoryAndClean(s string, previous string) string {

	_, errTime := parseTimeAllFormats(s)

	switch {
	case s == "---":
		return "delimiter:"
	case strings.HasPrefix(s, "log") && strings.HasSuffix(s, "txt"):
		return "name:" + s
	case s == "TIME1" || s == "TIME2" || s == "CSV" || s == "JSON" || s == "CONSOLE":
		return s
	case s == "Результат:" || s == "":
		return ""
	case strings.Contains(previous, "description:") || strings.Contains(previous, "log:"):
		return "log:" + s
	case errTime == nil && strings.Contains(previous, "time:") || strings.Contains(previous, "TIME"):
		return "time:" + s
	case strings.Contains(previous, "console:") || strings.Contains(previous, "CONSOLE"):
		return "console:" + s
	case strings.Contains(previous, "csv:") || strings.Contains(previous, "CSV"):
		return "csv:" + s
	case strings.Contains(previous, "JSON"):
		return "json:" + s
	default:
		return "description:" + s
	}
}

func ParseTests(path string, rewrite bool) []TestCase {
	filesData := analyzeExamplesTXT(path)
	res := []TestCase{}
	for _, fileData := range filesData {

		name := fileData[0][0]
		descr := fileData[1][0]
		logs := fileData[2]
		if logs[0] == "(файл пуст)" {
			logs = []string{}
		}

		time1, err := parseTimeAllFormats(fileData[3][0])
		if err != nil {
			log.Fatalf("Неверный формат времени в \"examples.txt\"")
			return []TestCase{}
		}

		time2, err := parseTimeAllFormats(fileData[3][1])
		if err != nil {
			log.Fatalf("Неверный формат времени в \"examples.txt\"")
			return []TestCase{}
		}

		// var console []string
		// if fileData[4][0] == "(empty)" && len(fileData[4]) == 1 {
		// 	console = []string{""}
		// } else {
		// 	console = fileData[4]
		// }

		console := fileData[4]
		csv := fileData[5]
		if csv[0] == "(файл не создаётся)" {
			csv = []string{}
		}
		json := fileData[6][0]
		if json == "(файл не создаётся)" {
			json = ""
		}

		if rewrite {
			dump := []string{}
			dump = append(dump, fileData[3][0])
			dump = append(dump, fileData[3][1])
			dump = append(dump, logs...)
			writeFile("./testData/"+name, dump)
		}

		res = append(res, TestCase{
			Name:    name,
			Descr:   descr,
			Logs:    logs,
			Time1:   time1,
			Time2:   time2,
			Console: console,
			CSV:     csv,
			JSON:    json,
		},
		)
	}
	return res
}

func writeFile(s string, c []string) {

	// Открыть файл с правами на чтение и запись.
	// Создать, если не существует. Очистить при открытии (O_TRUNC).
	f, err := os.OpenFile(s, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Невозможно открыть или создать файл %s: %v\n", s, err)
	}
	defer f.Close()

	for _, v := range c {
		fmt.Fprintf(f, "%s\n", v)
	}
}

func analyzeExamplesTXT(path string) [][][]string {

	s, err := ReadTXTtoString(path)
	if err != nil {
		log.Fatalf("Проблемы при открытии \"example.txt\".")
	}

	r := []string{}
	previous := ""
	for _, str := range s {
		row := setCategoryAndClean(str, previous)
		if row != "" && row != "TIME1" && row != "TIME2" && row != "CSV" && row != "JSON" && row != "CONSOLE" && row != "delimiter:" {
			r = append(r, row)
		}
		previous = row
	}

	result := cutSlice(r)

	return result
}
