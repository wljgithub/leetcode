// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// 
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// 
// Example:
// 
// Given nums = [2, 7, 11, 15], target = 9,
// 
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].


// 暴力破解
// 最好情况时间复杂度是 O(1),
// 最坏时间复杂度，O(n^2)

func twoSum(nums []int, target int) []int {
    length:=len(nums)
    for i:=0;i<length;i++{
        for j:=0;j<length;j++{
            if i!=j&& nums[i]+nums[j]==target{
                return []int{i,j}
            }
        }

    }
    return []int{}
}

// 用哈希表，空间换时间
// 时间复杂度 O(n)

func twoSum(nums []int, target int) []int {
    if len(nums)<2{
        return []int{}
    }
    hashMap:=make(map[int]int)

    for i,v:=range nums{
        hashMap[v] = i
    }

    var diff int
    for i,v:=range nums{
        diff = target - v
        if index,ok:=hashMap[diff];ok&&i!=index{
            return []int{i,index}
        }
    }

    return []int{}
}

