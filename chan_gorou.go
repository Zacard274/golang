package main

import (
	"fmt"
	"time"
)

func main() {
	var i = 3
	go func(a int) {
		fmt.Println(a)
	}(i)
	fmt.Println("1")
	time.Sleep(1 * time.Second)
}
