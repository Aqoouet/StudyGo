package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type StringGrouper struct {
	Str         string
	MaxPerGroup int
	GroupData   map[string][][]string
	DataJson    string
	DataText    string
}

func main() {
	groups := &StringGrouper{Str: "apple, Apricot, Avocado, banana, Blueberry",
		MaxPerGroup: 2,
	}
	groups = groups.Handle()
	if groups != nil {
		fmt.Printf("%v\n", groups.DataText)
	} else {
		fmt.Printf("Программа прервана из-за ошибок\n")
	}
}

func (groups *StringGrouper) Handle() *StringGrouper {
	err := groups.Parse()
	if err == nil {
		groups.DataJson = groups.AddJSON()
		groups.DataText = groups.AddTextData()
		return groups
	}
	return nil
}

func (sg *StringGrouper) Parse() error {
	if sg.MaxPerGroup <= 0 {
		return fmt.Errorf("maxPerGroup должно быть > 0")
	}
	if sg.Str == "" {
		return fmt.Errorf("входная строка пуста")
	}

	slice := StrToSlice(sg.Str)
	if len(slice) == 0 {
		return fmt.Errorf("все строки оказались пустыми")
	}

	cut := CutSlice(slice, sg.MaxPerGroup)
	sg.GroupData = Group(cut)
	return nil
}

func StrToSlice(s string) []string {
	if s == "" {
		return []string{}
	}
	clean := strings.Replace(s, " ", "", -1)
	slice := strings.Split(clean, ",")
	lower := [][]string{}
	for _, val := range slice {
		if val == "" {
			continue
		}
		symbol := ChangeSymbol([]rune(val)[0])
		tmp := []string{val, strings.ToLower(string(symbol))}
		lower = append(lower, tmp)
	}
	sort.SliceStable(lower, func(i, j int) bool { return lower[i][1] < lower[j][1] })
	res := []string{}
	for i, _ := range lower {
		res = append(res, lower[i][0])
	}
	return res
}

func CutSlice(s []string, n int) [][]string {
	if len(s) == 0 || n == 0 {
		return [][]string{}
	}
	rez := [][]string{}
	slice := []string{}
	for _, val := range s {
		if len(slice) == n {
			rez = append(rez, slice)
			slice = []string{}
		}
		if len(slice) == 0 {
			slice = append(slice, val)
			continue
		}
		symbol0 := ChangeSymbol([]rune(strings.ToLower(slice[0]))[0])
		symbol := ChangeSymbol([]rune(strings.ToLower(val))[0])
		if symbol0 != symbol {
			rez = append(rez, slice)
			slice = []string{}
			slice = append(slice, val)
		} else {
			slice = append(slice, val)
		}
	}
	if slice != nil {
		rez = append(rez, slice)
	}
	return rez
}

func Group(cut [][]string) map[string][][]string {
	data := map[string][][]string{}
	var symbol string
	for _, val := range cut {
		symbol = string(ChangeSymbol([]rune(val[0])[0]))
		symbol = strings.ToLower(symbol)
		tmp, exist := data[symbol]
		if !exist {
			tmp = [][]string{val}
		} else {
			tmp = append(tmp, val)
		}
		data[symbol] = tmp
	}
	return data
}

func ChangeSymbol(a rune) rune {
	if !unicode.IsLetter(a) {
		return 63
	}
	return a
}

func (input *StringGrouper) AddJSON() string {
	jsonData, err := json.MarshalIndent(input.GroupData, "", "  ")
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		os.Exit(1)
	}
	return string(json.RawMessage(jsonData))
}

func (input *StringGrouper) AddTextData() string {
	var res strings.Builder
	keys := []string{}
	for key, _ := range input.GroupData {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var val [][]string
	for j, key := range keys {
		val, _ = input.GroupData[key]
		res.WriteString(fmt.Sprintf("\"%s\":\n", key))
		for ind, slice := range val {
			res.WriteString(fmt.Sprintf("  [%d] %s", ind+1, strings.Join(slice, ", ")))
			if ind < len(val)-1 {
				res.WriteString("\n")
			}
		}
		if j < len(keys)-1 {
			res.WriteString("\n")
		}
	}
	return res.String()
}
