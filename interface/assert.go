package main

import "fmt"

type I interface{}

type II interface {
	f1()
}

func g2(a II) {
	fmt.Println("xxx")
}

func h3(b I) {
	g2(b.(II))
}

type Message struct {
	Name string `json:"msg_name"`
}

func (*Message) f1() {
}

func main() {
	var m = Message{
		Name: "Alice",
	}
	h3(&m)
}
