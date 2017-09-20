// Comb Sort in Golang
package main

import (
	"fmt"
	"useful_code/data-structure-and-algorithms/util"
)

func main() {

	slice := util.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	combsort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

func combsort(items []int) {
	var (
		n       = len(items)
		gap     = len(items)
		shrink  = 1.3
		swapped = true
	)

	for swapped {
		swapped = false
		gap = int(float64(gap) / shrink)
		if gap < 1 {
			gap = 1
		}
		for i := 0; i+gap < n; i++ {
			if items[i] > items[i+gap] {
				items[i+gap], items[i] = items[i], items[i+gap]
				swapped = true
			}
		}
	}
}
