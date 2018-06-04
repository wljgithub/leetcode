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
        #用桶排序的方法能解决，先将相同的数都放进一个桶里面，然后将桶的下标设置为桶里的个数
        #再将桶按下标排序，然后从后往前数，即可

        #put the element into  bucket
        d = {}
        for i in nums:
        	if i in d:
        		d[i] += 1
        	else:
        		d[i] = 1

        # print(d)
      
        new_d=sorted(zip(d.keys(),d.values()),key= lambda x : x[1])
        # print(new_d)

        return [new_d[-i][0] for i in range(1,k+1)]


o=Solution()
print(o.topKFrequent([1,1,1,2,2,3],2))