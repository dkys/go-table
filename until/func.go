package until

import (
	"fmt"
	"unicode"
)

// Length 返回字符串在屏幕上占用位置
func Length(str string) int {
	length := 0
	for _, char := range str {
		switch {
		case unicode.Is(unicode.Han, char): // 判断是否为中文
			if char == '·' {
				fmt.Printf(">>>>>>%v\n", "·")
			}
			length += 2
		case unicode.Is(unicode.Common, char) && len(string(char)) > 2: // 判断是否为中文符号
			length += 2
		default:
			length++
		}
	}
	return length
}

func ToString(val any) string {
	return fmt.Sprintf("%v", val)
}
