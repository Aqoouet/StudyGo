package main

import (
    "fmt"
    "time"
)

// формат дней рождений
const layout = "02.01.2006"

func Age(birthday string) (int, error) {
	bd , err := time.Parse(layout,birthday)
	if err!=nil {
		return 0, nil
	}
	return int(time.Since(bd).Hours()/24/365), nil
}

func main() {
    for _, v := range []string{"04.01.1969", "28.07.1995",
        "16.12.2007", "07.03.1947"} {
        age, err := Age(v)
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Println(v, ":", age)
    }
}