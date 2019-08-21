/*
Write a function that takes a string as input and returns the string reversed.

Example:
Given s = "hello", return "olleh".

Seen this question in a real interview before?  
*/

func reverseString(s string) string {
	l := len(s)
	chars := []rune(s)
	for i,c := range s {
		chars[l-i-1] = c
	}
	return string(chars)
}