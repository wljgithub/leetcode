package main

import (
	"fmt"
	"math"
)

func reveserString(s string) string {
	chars := []rune(s)
	l := len(s)
	for i, v := range s {
		chars[l-1-i] = v
	}
	return string(chars)

}
func titleToNumber(s string) int {
	s = reveserString(s)
	var sum float64
	for i, v := range s {
		sum += float64((int(v) - 64)) * math.Pow(float64(26), float64(i))
		// fmt.Println(i, sum)
	}
	return int(sum)
}

func main() {
	fmt.Println(titleToNumber("ZY"))
	// fmt.Println(reveserString("ABCDEF"))
}
