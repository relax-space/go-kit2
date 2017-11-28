package random

import (
	"fmt"
	"testing"
)

func Test_UuId(t *testing.T) {
	uuIdInt := UuIdInt64()
	uuId := Uuid("")
	fmt.Println(uuIdInt)
	fmt.Println(uuId)
}

func Test_NewUuId(t *testing.T) {
	uuId := NewUuid("20")
	fmt.Println(uuId)
}
