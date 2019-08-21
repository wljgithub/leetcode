'''
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.
Example:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
'''

class Solution:
	'''
	使用双指针，一个指针从最小的元素开始，另一个从最大的元素开始
	如果两指针指向的元素 sum == target,则返回两元素的索引，如果sum < target,则移动较小的元素，如果sum > target,则移动较大的元素
	'''
	def twoSum(self, numbers, target):

		"""
		:type numbers: List[int]
		:type target: int
		:rtype: List[int]
		"""
		
		l,h=0,len(numbers)-1
		while l<h:
			print(l,h)
			sum = numbers[l]+numbers[h]
			if sum == target:
				return l+1,h+1
			elif sum < target:
				l+=1
			else:
				h-=1

o=Solution()
print(o.twoSum([2,7,11,15],9))