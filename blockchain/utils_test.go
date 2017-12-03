package blockchain

import (
	"testing"
	"fmt"
)

func TestIntToHex(t *testing.T) {
	var a int64 = 255
	b := IntToHex(a)
	if b == nil {
		t.Error("is error")
	}
	fmt.Println(b)


}