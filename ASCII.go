package main

import (
	"fmt"
)

func  main() {
	a :=[]string{"0","1","2","3","4","5","6","7","8","9","0"}
	for _,item :=range a {
		fmt.Printf("%d\n",[]rune(item))
	}
}
