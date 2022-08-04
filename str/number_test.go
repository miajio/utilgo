package str_test

import (
	"fmt"
	"testing"

	"github.com/miajio/utilgo/str"
)

func TestStr(t *testing.T) {
	fmt.Println(str.IndexOf("测试一波不会出事", "一波"))

	fmt.Println(str.GetNumber("192.365.123.456"))

	fmt.Println(str.SubString("测试一波不会出事", 0, 3))
}
