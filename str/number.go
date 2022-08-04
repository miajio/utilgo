package str

import (
	"strconv"
)

const (
	default_number = "0123456789"
)

// GetNumber
// 获取字符串中的数字
func GetNumber(val string) int {
	if val == "" {
		return 0
	}

	var result string
	for i := 0; i < len(val); i++ {
		for j := 0; j < len(default_number); j++ {
			if val[i] == default_number[j] {
				result += string(val[i])
				break
			}
		}
	}

	if result == "" {
		return 0
	}

	r, _ := strconv.Atoi(result)
	return r
}
