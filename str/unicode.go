package str

import (
	"fmt"
	"strings"
)

// Contains
// 判断val的字符串中是否包含base的字符集
func Contains(val, base string) bool {
	return IndexOf(val, base) >= 0
}

// IndexOf
// 获取val的字符串中base的第一次出现的位置
func IndexOf(val, base string) int {
	dex := strings.Index(val, base)
	fmt.Println(dex)
	if dex > 0 {
		prefix := []byte(val)[0:dex]
		dex = len([]rune(string(prefix)))
	}
	return dex
}

// SubString
// 获取val的字符串中指定start和end的子字符串
func SubString(val string, start, end int) string {
	if start < 0 {
		start = 0
	}

	rval := []rune(val)

	if end > len(rval) {
		end = len(rval)
	}

	if start <= 0 && end >= len(rval) {
		return val
	}

	var result string
	for i := start; i < end; i++ {
		result += string(rval[i])
	}
	return result
}
