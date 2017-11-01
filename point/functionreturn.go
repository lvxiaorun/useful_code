package main

import (
	"fmt"
	"reflect"
)

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: 0}
}

func (c *Circle) DummyMethod() {
	fmt.Println("Type of receiver:", reflect.TypeOf(c))
}

func main() {
	//m := NewCircle(1)
	NewCircle(1).DummyMethod() // Chaining function calls results in error.
	//m := NewCircle(1)
	//m.DummyMethod()
}
