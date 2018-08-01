package main

import (
	"fmt"
)

func isVowels(r rune)(bool) {
	for _,v := range "aeiouAEIOU"{
		if v == r{
			return  true
		}
	}
	return false
}
func reverseVowels(s string) string {
	l,h:=0,len(s)-1
	str := []rune(s)
	for l<h{
		if isVowels(str[l]) && isVowels(str[h]){
			str[l],str[h] =str[h],str[l]
			l++;h--
			//fmt.Println(string(str))
		}
		if !isVowels(str[l]) {
			l++
		}
		if !isVowels(str[h]){
		h--
	}}
	return string(str)
}
// fmt.Println(string(str))


func main() {
	//var a string = "hello"
	//fmt.Println([]rune(a))
	//fmt.Println([]byte("a")[0])
	//fmt.Println("a"[0])
	fmt.Println(reverseVowels("hello"))
}
