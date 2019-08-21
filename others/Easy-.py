'''
假设你有一个很长的花圃，里面有些是种的，有些不是。然而，鲜花不能在邻近的土地上种植——它们会争夺水，而且两者都会死亡。
给定一个花床（表示为一个包含0和1的数组，其中0表示空，1表示不是空），和一个数字n，如果n个新花可以在其中植入，而不违反不相邻的花规则。

Example 1:
Input: flowerbed = [1,0,0,0,1], n = 1
Output: True
Example 2:
Input: flowerbed = [1,0,0,0,1], n = 2
Output: False
'''

class Solution:
    def canPlaceFlowers(self, flowerbed, n):
        """
        :type flowerbed: List[int]
        :type n: int
        :rtype: bool
        
        time complexity:O(n)
        """
        i=1
        while i < len(flowerbed) :
            interval = abs(flowerbed[i] - flowerbed[i-1])
            if interval >= n:
                return True
            i+=1
        return False
        
o=Solution()
print(o.canPlaceFlowers([1,0,0,0,1],2))