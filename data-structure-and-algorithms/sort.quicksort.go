package main

import "fmt"

func quick_sort(arr []int) ([]int) {
	if len(arr) <= 1 {
		return arr
	}
	length := len(arr)
	//把第一个作为标志数
	cNum := arr[0]
	var left []int
	var right []int

	for i := 1; i < length; i++ {
		if cNum >= arr[i] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	//左边结果继续排序
	left = quick_sort(left)
	right = quick_sort(right)
	ret := append(append(left, cNum), right...)
	return ret
}

func main() {
	arr := []int{5, 4, 8, 34, 24, 6, 45, 67, 5, 99, 4, 1, 21, 34, 45, 6, 76}
	ret := quick_sort(arr)
	fmt.Println(ret)

}
