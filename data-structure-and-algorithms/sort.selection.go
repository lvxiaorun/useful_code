package main

import (
	"fmt"
	"useful_code/data-structure-and-algorithms/util"
)

func main() {

	slice := util.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	selectionsort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

func selectionsort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}
