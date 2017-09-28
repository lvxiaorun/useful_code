package main

import (
	"time"
	"sync"
)

var c = make(chan int,5	)
func main(){
	ws := sync.WaitGroup{}
	ws.Add(10)
	now := time.Now().Unix()
	for i:=0;i<10;i++{
		select {
		case c<- i:
			go func(i int) {
				time.Sleep(1*time.Second)
				<- c
				ws.Done()
			}(i)
		}
	}
	ws.Wait()
	println(time.Now().Unix()-now)
}