package permutation

import (
	"testing"
	"fmt"
)

func TestZuheResult(t *testing.T) {
	re := ZuheResult(3,2)
	fmt.Println(re)
}


func TestPailieResult(t *testing.T) {
	re := PailieResult(3,2)
	fmt.Println(re)
}

func TestQuanPailie(t *testing.T) {
	res := QuanPailie([]int{1,2,3})
	fmt.Println(res)
}