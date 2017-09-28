package main

import "time"

func main() {
	now := time.Now().Unix()
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
	}
	println(time.Now().Unix() - now)

}
