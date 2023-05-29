package main

import (
	"fmt"
	"sync"
)

func main() {
	writes := 1000
	//var storage map[int]int

	storage := make(map[int]int, writes)
	wg := sync.WaitGroup{}
	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() { //хэш таблицей нельзя пользоваться конкурентно
			defer wg.Done()
			storage[i] = i
		}()
	}

	wg.Wait()
	fmt.Println(storage)

}
