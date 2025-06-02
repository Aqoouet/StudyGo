package main

import (
	//"time"
	"fmt"
)


func main() {
	
	ch := make(chan string)

	go func () {

		ch <- "Test"
	}()

	msg:=<-ch

	fmt.Println(msg)

	}
	