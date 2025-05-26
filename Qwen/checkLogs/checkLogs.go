package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

// func (la *LogAnalyzer) GroupBySource() map[string][]LogEntry
// func (la *LogAnalyzer) SaveToFile(filename, format string) error

var (
	ErrCritLowStringNumber = errors.New("критическая ошибка: недостаточно строк во входном файле")
	ErrCritTime            = errors.New("критическая ошибка при парсинге времени")
	ErrCritInputFile       = errors.New("критическая ошибка при открытии входного файла")
	ErrCritEmptyLogList    = errors.New("критическая ошибка: пустой список логов")

	ErrTime                 = errors.New("ошибка при парсинге времени")
	ErrWrongLogStringFormat = errors.New("ошибка: неверный формат строки")
)

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Source    string
	Message   string
	ToStat    bool
	//ErrStr    error
}

type Statistic struct {
	print      string
	printArray []string
	data       map[string]int
}

type LogData struct {
	nonFiltered []LogEntry
	filtered    []LogEntry
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

	var rez []LogEntry
	var errs []error
	for i := 2; i < len(s); i++ {
		v := s[i]
		le, err := handleStringData(v)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		rez = append(rez, le)
	}

	if rez == nil {
		return &emptyAnalyzer, []error{fmt.Errorf("нет логов удовлетворяющих формату: %w", ErrCritEmptyLogList)}
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

	row := strings.Split(s, "|")
	for i := range row {
		row[i] = strings.TrimSpace(row[i])
	}

	if len(row) <= 3 {
		return empty, fmt.Errorf("строка %q пропущена — %w", s, ErrWrongLogStringFormat)
	}

	for _, v := range row {
		if len(v) == 0 {
			return empty, fmt.Errorf("строка %q пропущена — %w", s, ErrWrongLogStringFormat)
		}
	}

	st := true
	t, err := parseTimeAllFormats(strings.TrimSpace(row[0]))
	if err != nil {
		st = false
	}
	l := strings.TrimSpace(row[1])
	so := strings.TrimSpace(row[2])
	m := strings.TrimSpace(row[3])

	le := LogEntry{
		Timestamp: t,
		Level:     l,
		Source:    so,
		Message:   m,
		ToStat:    st,
	}

	return le, nil

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

func getConsole(path string) []string {

	analyzer, errs := NewLogAnalyzer(path)

	for _, e := range errs {
		if errors.Is(e, ErrCritInputFile) || errors.Is(e, ErrCritLowStringNumber) || errors.Is(e, ErrCritTime) || errors.Is(e, ErrCritEmptyLogList) {
			return []string{e.Error()}
		}
	}
	analyzer.logs.filtered, _ = analyzer.FilterByPeriod(analyzer.start, analyzer.end)
	analyzer.Stat.data = analyzer.CountByLevel()
	analyzer.Stat.printArray, analyzer.Stat.print = analyzer.arrangeMessage(errs)
	// analyzer.PrintStats()

	return analyzer.Stat.printArray
}

func main() {
	getConsole("testData/log6.txt")
}
