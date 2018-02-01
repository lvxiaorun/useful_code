package main

import "fmt"

func ExFunc2(n int) func() {
	return func() {
		n++ // 这里对外部变量加 1
		fmt.Println(n)
	}
}

func main() {
	myFunc := ExFunc2(10)
	myFunc() // 这里输出 11

	myAnotherFunc := ExFunc2(20)
	myAnotherFunc() // 这里输出 21

	myFunc()        // 这里输出 12
	myAnotherFunc() // 这里输出 22
}
