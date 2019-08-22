// 53. Maximum Subarray
// Easy
// 
// 4866
// 
// 183
// 
// Favorite
// 
// Share
// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
// 
// Example:
// 
// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.

func maxSubArray(nums []int) int {
    max,sum:=nums[0],nums[0]
    
    for _,v:=range nums[1:]{
        if sum<0{
            sum = v
        }else{
            sum+=v
        }
        if max<sum{
            max = sum
        }
    }
    return max
}
