package encrypting

import (
	"testing"
	"fmt"
)

var key = []byte("lvlvztzt")

func TestDesEncrypt(t *testing.T) {
	ord := []byte("love")
	fmt.Println(ord)
	jia, err := DesEncrypt(ord, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jia)
	fmt.Println(string(jia))
	jie, _ := DesDecrypt(jia, key)
	fmt.Println(string(jie))
}
