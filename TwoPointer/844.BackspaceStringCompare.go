package main

import "fmt"

func backspaceCompare(S string, T string) bool {
	Ssum,Tsum := 0,0
	x,y := 0,0

	for i:=len(S)-1;i>0;i--{
		if S[i] ==byte(35){
			x ++
		}else if x>0{
			x --
		}else {
			Ssum += int(S[i])
		}
	}
	for i:=len(T)-1;i>0;i--{
		if T[i] ==byte(35){
			y ++
		}else if y>0{
			y --
		}else {
			Tsum += int(T[i])
		}
	}
	fmt.Println(Ssum,Tsum)

	return  Ssum ==Tsum
}

func main() {
	backspaceCompare("a#b#ab","abc#",)
	backspaceCompare("##","##a#")
	backspaceCompare("ab##","c#d#")
	//for _,k :=range "#"{
	//	fmt.Println(k,k == rune(35))
	//
	//}

}
