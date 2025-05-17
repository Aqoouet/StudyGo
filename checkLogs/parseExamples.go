package main

import (
	"bufio"
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

var TestCases = []TestCase{}

func containsTime(s string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", s)
	return err == nil
}

func readTXT(path string) [][][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Невозможно прочитать файл %s", path)
	}
	s := bufio.NewScanner(f)
	r := []string{}
	previous := ""
	for s.Scan() {
		row := setCategoryAndClean(s.Text(), previous)
		if row != "" && row != "TIME1" && row != "TIME2" && row != "CSV" && row != "JSON" && row != "CONSOLE" && row != "delimiter:" {
			r = append(r, row)
		}
		previous = row
	}
	if s.Err() != nil {
		log.Fatalf("Невозможно при чтении из файла %s", path)
	}
	return cutSlice(r)
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

	// name
	// description
	// log
	// time
	// console
	// csv
	// json

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
	case containsTime(s) && strings.Contains(previous, "time:") || strings.Contains(previous, "TIME"):
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

func parseTests(path string) []TestCase {
	filesData := readTXT(path)
	res := []TestCase{}
	for _, fileData := range filesData {

		name := fileData[0][0]
		descr := fileData[1][0]
		logs := fileData[2]
		if logs[0] == "(файл пуст)" {
			logs = []string{}
		}
		time1 := parseTime(fileData[3][0])
		time2 := parseTime(fileData[3][1])
		console := fileData[4]
		csv := fileData[5]
		if csv[0] == "(файл не создаётся)" {
			csv = []string{}
		}
		json := fileData[6][0]
		if json == "(файл не создаётся)" {
			json = ""
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

func parseTime(s string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		log.Fatalf("Ошибка при конвертации времени.")
	}
	return t
}

func main() {
	t := parseTests("examples.txt")
	for i, v := range t {
		fmt.Printf("%d: %v\n", i, v.JSON)
	}

}
