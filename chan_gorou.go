package main

import (
	"fmt"
)

func main() {
	var i = 3
	ch := make(chan int)
	go func(a int) {
		fmt.Println(a)
		v := <-ch
		fmt.Println(v)
	}(i)
	ch <- 4
	fmt.Println("1")

}
