package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	str := []byte(s)
	left, right := 0, len(str)-1
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
	fmt.Println(string(str))
	if string(str) == s {
		return true
	}
	return false

}

func main() {
	fmt.Println(isPalindrome("abc"))
	fmt.Println('a'-'A', 'b'-'B')
}
