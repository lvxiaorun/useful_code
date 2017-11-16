package main

import "fmt"

type a interface {
	Say()
}

type One struct {
}

type Two struct {
}

func (one One) Say() {
	fmt.Println("one")
}

func (two Two) Say() {
	fmt.Println("two")
}

func main() {
	var b, c a
	b = One{}
	c = Two{}
	a.Say(b) //if use interface (a) direct ,you can transference a param -> a instance which realization this interface
	b.Say()
	c.Say()
}
