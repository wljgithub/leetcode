package main

import (
	"fmt"
	"strconv"
)

func addDigits(num int) int {
	str := strconv.Itoa(num)
	for {
		sum := 0
		l := len(str)
		if l > 1 {
			for _, v := range str {
				sum += (int(v) - 48)
			}
			str = strconv.Itoa(sum)
		} else {
			v, _ := strconv.Atoi(str)
			return v
		}
	}
}

func main() {
	fmt.Println(addDigits(18))

}
