// Merge Sort in Golang
package main

import (
	"fmt"
	"useful_code/data-structure-and-algorithms/util"
)

func main() {

	slice := util.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	fmt.Println("\n--- Sorted ---\n\n", mergeSort(slice), "\n")
}

func mergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right)) //递归
}

func merge(left, right []int) (result []int) {
	fmt.Println(left)
	fmt.Println(right) //算法此函数得到的参数(数组)至少保证自己有序
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return
}
