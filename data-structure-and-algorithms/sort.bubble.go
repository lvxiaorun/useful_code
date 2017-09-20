// Bubble Sort in Golang
package main

import (
	"fmt"
	"useful_code/data-structure-and-algorithms/util"
)

func main() {

	slice := util.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubblesort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

func bubblesort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
