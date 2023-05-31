// Задача на телепликациююю Достаём уникальные числа
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	// 2 Создаём хэш таблицу, куда можем сложить только 1 уникальный ключ
	alreadyStored := make(map[int]struct{}) // множество с уникальными ключами(значение не требуется)
	mu := sync.Mutex{}
	capacity := 1000

	// Слайс длинной 1000 с двойниками, с 10 уникальными значениями
	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10)) // create rand num 0...9
	}

	// 1 Создаём канал, куда записываем уникальные значения
	uniqueIDs := make(chan int, capacity)
	wg := sync.WaitGroup{}

	for i := 0; i < capacity; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock() // блокировку ставим даже при чтении, так как возможны задвоения в записи при одновременном входе
			defer mu.Unlock()
			if _, ok := alreadyStored[doubles[i]]; !ok { //читаем есть ли ключ в мапе
				alreadyStored[doubles[i]] = struct{}{} // записываем ключ, если в мапе не нашли

				uniqueIDs <- doubles[i]
			}
		}()
	}

	wg.Wait() // ждём пока отработают гор

	// утины
	close(uniqueIDs) //закрываем канал иначе будет дедлок
	for val := range uniqueIDs {
		fmt.Println(val)
	}

	fmt.Printf("len of ids: %d\n", len(uniqueIDs)) // 0, 1, 2, 3, 4 ...
	fmt.Println(uniqueIDs)

}
