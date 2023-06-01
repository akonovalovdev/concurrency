package main

import "fmt"

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4
	close(ch1)
	close(ch2)

	ch3 := syncMerge[int](ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}
}

//func merge[T any](chans ...chan T) chan T {
//	return nil
//}

func syncMerge[T any](chans ...chan T) chan T { // синхронная функция
	l := 0
	for _, singleCh := range chans {
		l += len(singleCh)
	}
	result := make(chan T, l)
	for _, singleCh := range chans {
		for val := range singleCh {
			result <- val
		}
	}
	close(result)

	return result
}
