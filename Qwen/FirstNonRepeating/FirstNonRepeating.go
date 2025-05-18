package main

// Write a function named first_non_repeating_letter† that takes a string input, and returns the first character that is not repeated anywhere in the string.

// For example, if given the input 'stress', the function should return 't', since the letter t only occurs once in the string, and occurs first in the string.

// As an added challenge, upper- and lowercase letters are considered the same character, but the function should return the correct case for the initial letter. For example, the input 'sTreSS' should return 'T'.

// If a string contains all repeating characters, it should return an empty string ("");

// † Note: the function is called firstNonRepeatingLetter for historical reasons, but your function should handle any Unicode character.

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		iStr := string(runes[i])
		jStr := string(runes[j])
		iStrLower := strings.ToLower(iStr)
		jStrLower := strings.ToLower(jStr)
		iLower := rune(iStrLower[0])
		jLower := rune(jStrLower[0])
		return iLower < jLower
	})

	result := string(runes)

	return result
}

func deleteElm(s []rune, del rune) []rune {
	cleaned := []rune{}
	for _, j := range s {
		if j != del {
			cleaned = append(cleaned, j)
		}
	}
	return cleaned
}

func FirstNonRepeating(str string) string {

	if len(str) == 1 || len(str) == 0 {
		return str
	}

	sortedString := sortString(str)
	sortedStringLower := strings.ToLower(sortedString)
	runes := []rune(sortedStringLower)
	runesInitial := []rune(str)

	symbolsNotRepeated := []rune{}
	symbolsRepeated := []rune{}
	resultIndex := 0

	for i := 0; len(runes) != 0; i++ {
		i = 0
		j := runes[i]

		case1 := (len(runes) == 1 || i == 0 && runes[0] != runes[1])
		case2 := (len(runes) != 1 && i == len(runes)-1 && runes[len(runes)-1] != runes[len(runes)-2])
		case3 := (len(runes) != 1 && i != 0 && i != len(runes)-1 && runes[i-1] != runes[i] && runes[i+1] != runes[i])

		if case1 || case2 || case3 {
			symbolsNotRepeated = append(symbolsNotRepeated, j)
		} else {
			symbolsRepeated = append(symbolsRepeated, j)
		}

		if len(runes) != 0 {
			runes = deleteElm(runes, j)
		}

	}

	if len(symbolsNotRepeated) == 0 {
		return ""
	}

finish:
	for i, j := range runesInitial {
		jStr := string(j)
		jStrLower := strings.ToLower(jStr)
		jLower := rune(jStrLower[0])
		for _, b := range symbolsNotRepeated {
			if b == jLower {
				resultIndex = i
				break finish
			}
		}
	}

	return string(runesInitial[resultIndex])
}

func main() {

	str := ""
	str = "a"
	fmt.Println(str, "==>", FirstNonRepeating(str), "==>", "a")
	str = "stress"
	fmt.Println(str, "==>", FirstNonRepeating(str), "==>", "t")
	str = "moonmen"
	fmt.Println(str, "==>", FirstNonRepeating(str), "==>", "e")
	str = ""
	fmt.Println(str, "==>", FirstNonRepeating(str), "==>", "")
	str = "abba"
	fmt.Println(str, "==>", FirstNonRepeating(str), "==>", "")
	str = "aa"
	fmt.Println(str, "==>", FirstNonRepeating(str), "==>", "")
	str = "~><#~><"
	fmt.Println(str, "==>", FirstNonRepeating(str), "==>", "#")
	str = "hello world, eh?"
	fmt.Println(str, "==>", FirstNonRepeating(str), "==>", "w")
	str = "sTreSS"
	fmt.Println(str, "==>", FirstNonRepeating(str), "==>", "T")
	str = "Go hang a salami, I'm a lasagna hog!"
	fmt.Println(str, "==>", FirstNonRepeating(str), "==>", ",")
}
