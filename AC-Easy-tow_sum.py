'''
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
'''


#嵌套一个循环，双重遍历，时间复杂度超时
class Solution:
	'''
	时间复杂度:O(n**2)
	'''
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for index_1,i in enumerate(nums):
            for index_2,j in enumerate(nums):
                if i+j==target and index_1!=index_2:
                    return [index_1,index_2]

#成功
class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for i,v in enumerate(nums):
            value=target-v
            if value in nums and i!=nums.index(value):
                return sorted([i,nums.index(value)])