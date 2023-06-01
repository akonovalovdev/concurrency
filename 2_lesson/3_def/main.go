package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case val := <-ch:
		fmt.Println(val)
	}
}
