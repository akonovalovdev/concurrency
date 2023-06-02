package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4
	close(ch1)
	close(ch2)

	ch3 := merge[int](ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}
}

func merge[T any](chans ...chan T) chan T {
	result := make(chan T)
	wg := sync.WaitGroup{} //wait group когда все горутины запишут, канал должен закрыться
	for _, singleChan := range chans {
		wg.Add(1) // давайте добавим задачку
		singleChan := singleChan
		go func() {
			defer wg.Done() // после того как горутина отработала - задачку снимаем
			for val := range singleChan {
				result <- val
			}
		}()
	}

	// так как из канала никто не читает, а он создан без буфера wg.Wait необходимо вызывать в отдельной горутине,
	//а не в мэйне

	return result
}

//func syncMerge[T any](chans ...chan T) chan T { // синхронная функция
//	l := 0
//	for _, singleCh := range chans {
//		l += len(singleCh)
//	}
//	result := make(chan T, l) // буфер необходим так как функция синхронная
//	for _, singleCh := range chans {
//		for val := range singleCh {
//			result <- val
//		}
//	}
//	close(result)
//
//	return result
//}
