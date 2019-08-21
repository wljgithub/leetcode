package main

import "fmt"

func singleNumber(nums []int) []int {
	var e0,e1 int
	for _,v:=range nums{
		e0^=v
	}
	one:=e0&-e0
	for _,v :=range nums{
		if one&v>0{
			e1^=v
		}
	}
	return []int{e1,e0^e1}
}

func main() {
	var arr = []int{1,2,1,3,2,5}
	fmt.Println(singleNumber(arr))
	var i int32 = -121
	fmt.Println(^i)
	//var unsign uint8 =1
	//var sign int32 =0x0fffffff

	fmt.Println(3/2)
	var u uint8 =1
	fmt.Println(^u)
	fmt.Println(arr[:1])
}
