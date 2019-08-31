// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
// 
// Note:
// 
// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
// 
// Example 1:
// 
// Input: [2,2,1]
// Output: 1
// Example 2:
// 
// Input: [4,1,2,1,2]
// Output: 4


// 利用异或的特性，相同的数异或会抵消掉，比如3^3^2等于2
func singleNumber(nums []int) int {
    if len(nums)<1{
        return -1
    }
    num:=nums[0]
    
    for _,v:=range nums[1:]{
        num^=v
    }
    
    return num
}
