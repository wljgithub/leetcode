package main

import (
		"fmt"
	"strconv"
)

func addDigits(num int) int {
	l:=len(strconv.Itoa(num))

	//fmt.Println(str,l,sum)
	for  l>1{
		sum :=0
			for _,k:=range strconv.Itoa(num){
				value,err:= strconv.Atoi(string(k))
				if err==nil{

					sum += value
				}
			}
			num = sum
			l=len(strconv.Itoa(num))
			fmt.Println("sum",sum,"len:",l)

		}
	return num
	}


func main()  {
	//fmt.Println(len("12"))

	//fmt.Println(strconv.Itoa(1))

	fmt.Println(addDigits(66))
	//str := "123"
	//sum := 0
	//for _,k:=range str{
	//	//fmt.Println(string(k))
	//	value,err:= strconv.Atoi(string(k))
	//	if err ==nil{
	//		sum += value
	//	}
	//}
	//fmt.Println(sum)


}
