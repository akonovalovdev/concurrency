// задача про тайм аут
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	chanForResp := make(chan int) // канал который будет получать ответ
	go RPCCall(chanForResp)       //вызов удалённой процедуры

	result := <-chanForResp
	fmt.Println(result)
}

func RPCCall(ch chan<- int) { //может быть сеть или сервер
	// sleep 0-4 sec
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))

	ch <- rand.Int()
}
