// Требования:
// Не использовать регулярные выражения.
// Не использовать сторонние библиотеки.
// Реализовать всё "вручную", через циклы и условия.
// Обработать краевые случаи:
// - Пустая строка
// - Строка без повторений (abcd)
// - Строка, где все символы одинаковые (aaaaa)
// - Некорректный формат при распаковке

// func Compress:
// Каждую последовательность одинаковых подряд идущих символов
// заменить на: cимвол + количество, если длина серии > 1
// просто символ, если серия из одного символа
//
// Если результат сжатия не короче исходной строки —
// вернуть оригинальную строку.

// func Decompress:
// Восстанавливает оригинальную строку из сжатой.
// Должна обрабатывать ошибки (например,
// некорректный формат входной строки).

// IsValidCompressedFormat:
// функция проверяет является ли строка сжимаемой

package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Decompress("12a3"))
}

func Decompress(s string) string {
	if s == "" {
		return ""
	}
	b := []rune(s)
	var res strings.Builder
	_, initI := findSubString(b, 0)
	if initI != 0 {
		res.WriteString(string(b[:initI]))
	}
	for i := initI; i < len(b); {
		char := string(b[i])
		if i+1 == len(b) {
			res.WriteString(char)
			return res.String()
		}
		_, endDigit := findSubString(b, i)
		if endDigit == 0 && !isDigit(b[i+1]) {
			res.WriteString(char)
			i++
			continue
		}
		_, endDigitNext := findSubString(b, i+1)
		subString := string(b[i+1 : endDigitNext])
		subDigit, _ := strconv.ParseInt(subString, 10, 64)
		new := string(repeatRune(b[i], int(subDigit)))
		res.WriteString(new)
		i += len(subString) + 1
	}
	return res.String()
}

type compData struct {
	InitialString string
	FinalString   string
}

// Если n соответствует цифре, то функция пытается определить где заканчивается число
// Если n соответствует не цифре, а символу, то функция пытается определить повторяется ли символ и где конец повторений
func findSubString(b []rune, n int) (int, int) {
	val := b[n]
	endVal := n
	for endVal < len(b) && b[endVal] == val {
		endVal++
	}
	if !isDigit(b[n]) {
		return endVal, 0
	}
	endDigit := n
	for endDigit < len(b) && isDigit(b[endDigit]) {
		endDigit++
	}
	return endVal, endDigit
}

func StrToBytesSlice(s string) [][]rune {
	b := []rune(s)
	bSliced := [][]rune{}
	for i := 0; i < len(b); {
		endVal, _ := findSubString(b, i)
		bSliced = append(bSliced, b[i:endVal])
		i = endVal
	}
	return bSliced
}

func isDigit(a rune) bool {
	if a >= '0' && a <= '9' {
		return true
	}
	return false
}

func repeatRune(a rune, n int) []rune {
	res := []rune{}
	for i := 0; i < n; i++ {
		res = append(res, a)
	}
	return res
}

func Compress(s string) string {
	if s == "" {
		return ""
	}
	var strFull strings.Builder
	var str string
	bsliced := StrToBytesSlice(s)
	for _, item := range bsliced {
		if len(item) == 1 {
			str = string(item[0])
		} else {
			str = string(item[0]) + string(strconv.Itoa(len(item)))
		}
		strFull.WriteString(str)
	}
	res := strFull.String()
	if len(res) == len(s) {
		return s
	} else {
		return res
	}
}

func ReplaceAtPos(runes, old, new []rune, n int) []rune {
	return slices.Concat(runes[:n], new, runes[n+len(old):])
}

func IsValidCompressedFormat(s string) bool {
	if s != Decompress(s) {
		return true
	}
	return false
}
