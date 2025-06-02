package main

import (
	"fmt"
	"strconv"
	"strings"
)

var  currency = map[string][2]string{
	"Маврикийская рупия":			{"U+20A8","0.021"},
	"Бутанский нгултрум":			{"U+0F54","0.013"},
	"Лаосский кип":					{"U+20AD","0.00009"},
	"Монгольский тугрик":			{"U+20AE","0.00034"},
}


func main() {
	fmt.Printf("%-20s\t%-20s\t%-20s\n","Название валюты","Символ Unicode","Курс к доллару")
	for key, currencyData := range currency{
		codeStr := 	strings.TrimPrefix(currencyData[0],"U+")
		symbol, _:= strconv.ParseInt(codeStr,16,32)
		exchageRate, _ := strconv.ParseFloat(currencyData[1],64)
		fmt.Printf("%-20s\t%-20c\t%-20.6f\n",key,symbol,exchageRate)
	}
}