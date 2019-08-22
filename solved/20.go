// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// 
// An input string is valid if:
// 
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.
// 
// Example 1:
// 
// Input: "()"
// Output: true
// Example 2:
// 
// Input: "()[]{}"
// Output: true
// Example 3:
// 
// Input: "(]"
// Output: false
// Example 4:
// 
// Input: "([)]"
// Output: false
// Example 5:
// 
// Input: "{[]}"
// Output: true


// 一开始这么写，没考虑到 ([)] 这种情况，然后没通过
// func isValid(s string) bool {
//     var(
//         parenthese int
//         bracket int
//         bigParenthese int
//     )
//     
//     for i:=0;i<len(s);i++{
//         switch s[i]{
//             case uint8('['):bracket++
//             case uint8(']'):bracket--
//             case uint8('('):parenthese++
//             case uint8(')'):parenthese--
//             case uint8('{'):bigParenthese++
//             case uint8('}'):bigParenthese--
//         }
//     }
//     
//     if !(parenthese==0&&bracket==0&&bigParenthese==0){
//         return false
//     }
//     return true
// }

// 用栈的方式
func isValid(s string) bool {
	stack:=[]byte{}

	for i:=0;i<len(s);i++{
		switch s[i]{
		case '(','[','{':
			stack = append(stack,s[i])
		case ')',']','}':
			if len(stack)==0||
				(s[i]==')'&&stack[len(stack)-1]!='(')||
				(s[i]==']'&&stack[len(stack)-1]!='[')||
				(s[i]=='}'&&stack[len(stack)-1]!='{'){
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0
}
