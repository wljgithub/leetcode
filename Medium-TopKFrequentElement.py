'''
Given a non-empty array of integers, return the k most frequent elements.

For example,
Given [1,1,1,2,2,3] and k = 2, return [1,2].

Note: 
You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
'''

class Solution:
    def topKFrequent(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """

        #put the element into  bucket
        # d = {}
        # for i in nums:
        # 	if i in d:
        # 		d[i] += 1
        # 	else:
        # 		d[i] = 1

        # print(d)
        n = set(nums)
        bucket = []
        for i in n :
        	bucket.append([i])
        print(bucket)




        #将排序后，找出元素最大的前K个桶
        #但是字典能根据值来排序的吗？排序后jian



o=Solution()
o.topKFrequent([1,1,1,2,2,3],2)