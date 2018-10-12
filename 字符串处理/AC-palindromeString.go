package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	str := []rune{}
	for _, v := range s {
		if (v >= 48 && v <= 57) || (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			str = append(str, v)
		}
	}
	fmt.Println(string(str))
	left, right := 0, len(str)-1
	for left < right {

		if ok := ignoreCaseSensitive(str[left], str[right]); !ok {
			return false
		}
		left++
		right--
	}

	return true

}

func ignoreCaseSensitive(a, b rune) bool {
	// 如果是大小字母，转为小写
	if a >= 65 && a <= 90 {
		a += 32
	}
	if b >= 65 && b <= 90 {
		b += 32
	}
	// 判断字符是否相等
	if a == b {
		return true
	}
	return false
}

func main() {
	// fmt.Println(isPalindrome("aA"))
	fmt.Println(isPalindrome("Zeus was deified, saw Suez."))

	// fmt.Println('a', 'z', 'A', 'Z', '0', '9')

}
