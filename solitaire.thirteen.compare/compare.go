package main

import (
	"useful_code/permutation"
	"fmt"
)

type Pai struct {
	Args  []int
	Times int
}

func main() {
	solitaire := []int{1, 2, 3}
	res := permutation.QuanPailie(solitaire)
	fmt.Println(res)
	pais := []Pai{}
	for _, te := range res {
		pais = append(pais, Pai{te, 0})
	}
	for bi, item := range pais {
		for _, item2 := range res {
			smalltime := 0
			for i, si := range item2 {
				if item.Args[i] > si {
					smalltime ++
				}
			}
			if smalltime >1 {
				pais[bi].Times ++
			}
		}
	}
	fmt.Println(pais)
}
