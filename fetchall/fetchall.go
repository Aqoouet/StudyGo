package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	calls = 2
	fName = "fetchResults.txt"
)

type Data struct {
	written  []int
	duration []float64
}

func main() {

	st := time.Now()
	maxC := len(os.Args[1:]) * calls
	ch := make(chan string)
	allData := make(map[string]*Data)

	f, err := os.Create(fName)
	defer f.Close()

	if err != nil {
		fmt.Printf("Ошибка создания файла %s: %v", fName, err)
		os.Exit(1)
	}

	for _, url := range os.Args[1:] {
		for call := 1; call <= calls; call++ {
			go fetch(url, ch)
		}
	}

	for read := 1; read <= maxC; read++ {
		readB := strings.Split(<-ch, ";")
		u := readB[0]
		bytes, _ := strconv.Atoi(readB[1])
		sec, _ := strconv.ParseFloat(readB[2], 64)
		if allData[u] == nil {
			allData[u] = &Data{}
		}
		allData[u].written = append(allData[u].written, bytes)
		allData[u].duration = append(allData[u].duration, sec)
	}

	fmt.Fprintf(f, "%30s%30v%30v\n", "Url", "Number of bytes", "Time to fetch")
	for k, d := range allData {
		sort.Ints(d.written)
		sort.Float64s(d.duration)
		fmt.Fprintf(f, "%30s%30v%30v\n", k, d.written[len(d.written)-1], formatSlice(d.duration))
	}

	et := time.Since(st).Seconds()
	fmt.Printf("Elapsed time to fetch: %.2fs.\n", et)
}

func fetch(url string, ch chan<- string) {
	stl := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Ошибка открытия url (%s): %v", url, err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("Ошибка копирования из url (%s): %v", url, err)
		return
	}
	dur := time.Since(stl).Seconds()
	ch <- fmt.Sprintf("%s;%d;%.2f", url, nbytes, dur)
}

func formatSlice[T int | float64](s []T) string {
	return fmt.Sprintf("%v", s)
}
