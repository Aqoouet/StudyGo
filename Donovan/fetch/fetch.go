package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"strings"
)

const (
	httpPrefix = "http://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix){
			url = httpPrefix + url
		}
		resp, err := http.Get(url)
		if err!= nil {
			fmt.Fprintf(os.Stderr, "ошибка открытия сайта: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf ("Статус открытия url: %v\n", resp.Status)
		siteContent, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err!=nil {
			fmt.Fprintf(os.Stderr, "ошибка чтения: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf ("%s",siteContent)
	}
}