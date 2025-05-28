package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	logID = 18
)

func rowCSV(k LogEntry) string {
	t := k.Timestamp.Format(TimeFormats[0])
	m := strings.Join([]string{t, k.Level, k.Source, k.Message}, ",")
	return m
}

func (la *LogAnalyzer) SortByTime() {

	s := []LogEntry{}
	s = append(s, la.logs.filtered...)

	sort.Slice(s, func(i, j int) bool {
		return la.logs.filtered[i].Timestamp.Before(la.logs.filtered[j].Timestamp)
	})

	la.logs.filteredSorted = s

}

func main() {
	rez := getCSV("testData/log" + strconv.Itoa(logID) + ".txt")
	for _, v := range rez {
		fmt.Println(v)
	}
}

// func (la *LogAnalyzer) GroupBySource() map[string][]LogEntry

var (
	ErrCritLowStringNumber = errors.New("критическая ошибка: недостаточно строк во входном файле")
	ErrCritTime            = errors.New("критическая ошибка при парсинге времени")
	ErrCritInputFile       = errors.New("критическая ошибка при открытии входного файла")
	ErrCritEmptyLogList    = errors.New("критическая ошибка: пустой список логов")
	ErrCritWrongTimeRange  = errors.New("критическая ошибка: конечное время меньше начального")

	ErrTime                 = errors.New("ошибка при парсинге времени")
	ErrWrongLogStringFormat = errors.New("ошибка: неверный формат строки")
)

type LogEntry struct {
	TimeStr   string    `json:"timestamp"`
	Timestamp time.Time `json:"-"`
	Level     string    `json:"level"`
	Source    string    `json:"source"`
	Message   string    `json:"message"`
	ToStat    bool      `json:"-"`
}

type Statistic struct {
	print      string
	printArray []string
	data       map[string]int
}

type LogData struct {
	nonFiltered    []LogEntry
	filtered       []LogEntry
	filteredSorted []LogEntry
}

var empty = LogEntry{}
var emptyAnalyzer = LogAnalyzer{}

type LogAnalyzer struct {
	filePath string
	start    time.Time
	end      time.Time
	logs     LogData
	Stat     Statistic
}

func ReadTXTtoString(path string) ([]string, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла %s: %w", path, err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	rez := []string{}
	for s.Scan() {
		rez = append(rez, s.Text())
	}

	err = s.Err()
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении из файла %s: %w", path, err)
	}

	if len(rez) == 0 {
		return nil, fmt.Errorf("ошибка: файл %s пустой", path)
	}

	return rez, nil
}

var TimeFormats = []string{
	"2006-01-02 15:04:05",
	"2006-01-02T15:04:05Z07:00",
}

func parseTimeAllFormats(s string) (time.Time, error) {

	for _, f := range TimeFormats {

		t, err := time.Parse(f, s)
		if err == nil {
			return t, nil
		}

	}
	return time.Time{}, fmt.Errorf("ошибка: неверный формат времени %q: %w", s, ErrTime)

}

func NewLogAnalyzer(path string) (*LogAnalyzer, []error) {

	var err error
	s, err := ReadTXTtoString(path)
	if err != nil {
		return &emptyAnalyzer, []error{fmt.Errorf("невозможно открыть входной файл: %w (%w)", err, ErrCritInputFile), err, ErrCritInputFile}
	}

	if len(s) <= 2 {
		return &emptyAnalyzer, []error{ErrCritLowStringNumber}
	}

	time1, err := parseTimeAllFormats(s[0])
	if err != nil {
		return &emptyAnalyzer, []error{fmt.Errorf("неверный формат времени в \"examples.txt\": %s (%w)", s[0], ErrCritTime)}
	}

	time2, err := parseTimeAllFormats(s[1])
	if err != nil {
		return &emptyAnalyzer, []error{fmt.Errorf("неверный формат времени в \"examples.txt\": %s (%w)", s[1], ErrCritTime)}
	}

	if time2.Before(time1) {
		return &emptyAnalyzer, []error{fmt.Errorf("t2 < t1: (%w)", ErrCritWrongTimeRange)}
	}

	var rez []LogEntry
	var errs []error
	for i := 2; i < len(s); i++ {
		v := s[i]
		le, err := handleStringData(v)
		if err != nil {
			errs = append(errs, err)
		}
		if c := errors.Is(err, ErrWrongLogStringFormat); c {
			continue
		} else {
			rez = append(rez, le)
		}
	}

	if rez == nil {
		e := fmt.Errorf("нет логов удовлетворяющих формату: %w", ErrCritEmptyLogList)
		errs = append(errs, e)
		return &emptyAnalyzer, errs
	}

	return &LogAnalyzer{
		filePath: path,
		start:    time1,
		end:      time2,
		logs: LogData{
			nonFiltered: rez,
			filtered:    []LogEntry{},
		},
	}, errs
}

func (la *LogAnalyzer) FilterByPeriod(start, end time.Time) ([]LogEntry, error) {
	var rez []LogEntry
	for _, di := range la.logs.nonFiltered {
		t := di.Timestamp
		st := di.ToStat
		if st && (t.Before(end) || t.Equal(end)) && (t.After(start) || t.Equal(start)) {
			rez = append(rez, di)
		}

	}
	if len(rez) == 0 {
		return nil, fmt.Errorf("нет строк удовлетворяющих временному фильтру: %w", ErrCritEmptyLogList)
	}

	return rez, nil
}

func (la *LogAnalyzer) PrintStats() {
	fmt.Println(la.Stat.print)
}

func (la *LogAnalyzer) CountByLevel() map[string]int {
	rez := map[string]int{}
	for _, v := range la.logs.nonFiltered {
		_, exist := rez[v.Level]
		if !exist {
			rez[v.Level] = 0
		}
	}
	for _, v := range la.logs.filtered {
		rez[v.Level] += 1
	}
	return rez
}

func handleStringData(s string) (LogEntry, error) {

	row1 := strings.SplitN(s, "|", 4)
	row := []string{}

	for _, v := range row1 {
		r := strings.Split(v, "\t")
		row = append(row, r...)
	}

	for i := range row {
		row[i] = strings.TrimSpace(row[i])
	}

	if len(row) <= 3 {
		return empty, fmt.Errorf("строка %q пропущена (недостаточно полей в записи) — %w", s, ErrWrongLogStringFormat)
	}

	for _, v := range row {
		if len(v) == 0 {
			return empty, fmt.Errorf("строка %q пропущена (имеются пустые поля в записи) — %w", s, ErrWrongLogStringFormat)
		}
	}

	l := strings.TrimSpace(row[1])

	wordsN := len(strings.Split(l, " "))
	if wordsN > 1 {
		return empty, fmt.Errorf("в обозначении level %q более одного слова, строка пропущена — %w", l, ErrWrongLogStringFormat)
	}

	allLetter := true
	for _, v := range l {
		if !unicode.IsLetter(v) {
			allLetter = false
		}
	}
	allUpper := true
	if l != strings.ToUpper(l) {
		allUpper = false
	}
	if !allLetter || !allUpper {
		return empty, fmt.Errorf("в обозначении level %q должны быть только прописные буквы, строка пропущена — %w", l, ErrWrongLogStringFormat)
	}

	so := strings.TrimSpace(row[2])
	m := unquote(strings.TrimSpace(row[3]))

	st := true
	t, err := parseTimeAllFormats(strings.TrimSpace(row[0]))
	if err != nil {
		st = false
	}

	le := LogEntry{
		TimeStr:   t.Format(TimeFormats[0]),
		Timestamp: t,
		Level:     l,
		Source:    so,
		Message:   m,
		ToStat:    st,
	}

	if err != nil {
		return le, fmt.Errorf("строка %q не учитывается в статистике (неверный формат времени): %w", s, ErrTime)
	} else {
		return le, nil
	}

}

func unquote(s string) string {

	splitted := strings.Split(s, "\\n")
	for i, v := range splitted {
		vUnquoted, errQuote := strconv.Unquote(`"` + v + `"`)
		if errQuote == nil {
			splitted[i] = vUnquoted
		}

	}
	return strings.Join(splitted, "\\n")
}

func (la *LogAnalyzer) arrangeMessage(errs []error) ([]string, string) {
	order := []string{"INFO", "ERROR", "DEBUG", "WARN"}

	var rez []string

	for _, e := range errs {
		rez = append(rez, e.Error())
	}

	added := map[string]int{}

	for _, lev := range order {
		if _, exist := la.Stat.data[lev]; exist {
			rez = append(rez, fmt.Sprintf("%s: %d", lev, la.Stat.data[lev]))
			added[lev] = 0
		}
	}

	nonStandard := []string{}

	for key := range la.Stat.data {
		if _, exist := added[key]; !exist {
			nonStandard = append(nonStandard, key)
			added[key] = 0
		}
	}

	sort.Strings(nonStandard)

	for _, lev := range nonStandard {
		rez = append(rez, fmt.Sprintf("%s: %d", lev, la.Stat.data[lev]))
	}

	var rezStr string

	for i := 0; i < len(rez)-1; i++ {
		rezStr += rez[i] + "\n"
	}
	rezStr += rez[len(rez)-1]

	return rez, rezStr
}

func criticalErrCheck(errs []error) (bool, []string) {

	textErrs := []string{}
	for _, v := range errs {
		textErrs = append(textErrs, v.Error())
	}

	for _, e := range errs {
		c1 := !errors.Is(e, ErrTime)
		c2 := !errors.Is(e, ErrWrongLogStringFormat)
		if c1 && c2 {
			return true, textErrs
		}
	}

	return false, textErrs

}

func getConsole(path string) []string {

	analyzer, crit, errs, textErrs := getAnalyzer(path)

	if crit {
		return textErrs
	}

	analyzer.Stat.data = analyzer.CountByLevel()
	analyzer.Stat.printArray, analyzer.Stat.print = analyzer.arrangeMessage(errs)

	return analyzer.Stat.printArray
}

func getCSV(path string) []string {

	analyzer, _, _, _ := getAnalyzer(path)

	if analyzer.logs.filteredSorted == nil {
		return []string{}
	}

	rez := []string{"timestamp,level,source,message"}

	for _, v := range analyzer.logs.filteredSorted {
		rez = append(rez, rowCSV(v))
	}

	return rez

}

func getJSON(path string) string {

	analyzer, _, _, _ := getAnalyzer(path)

	if analyzer.logs.filteredSorted == nil {
		return ""
	}

	jsonData, _ := json.Marshal(analyzer.logs.filteredSorted)

	return string(jsonData)

}

func getAnalyzer(path string) (*LogAnalyzer, bool, []error, []string) {

	analyzer, errs := NewLogAnalyzer(path)

	crit, textErrs := criticalErrCheck(errs)

	if !crit {
		analyzer.logs.filtered, _ = analyzer.FilterByPeriod(analyzer.start, analyzer.end)
	} else {
		return analyzer, crit, errs, textErrs
	}

	if analyzer.logs.filtered != nil {
		analyzer.SortByTime()
	}

	return analyzer, crit, errs, textErrs

}
