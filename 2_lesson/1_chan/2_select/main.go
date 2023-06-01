// задача про тайм аут
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RPC struct {
	jobs chan int
	rand *rand.Rand
}

func (r *RPC) Call() {
	time.Sleep(time.Second * time.Duration(r.rand.Intn(5))) // sleep 0-4 sec

	r.jobs <- rand.Int()
}

func main() {
	rpc := &RPC{
		jobs: make(chan int),
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	go rpc.Call() //вызов удалённой процедуры

	result := <-rpc.jobs
	fmt.Println(result)
}
