'''
Given a non-negative integer c, your task is to decide whether there're two integers a and b such that a2 + b2 = c.

Example 1:
Input: 5
Output: True
Explanation: 1 * 1 + 2 * 2 = 5
Example 2:
Input: 3
Output: False
'''

class Solution:
    def judgeSquareSum(self, c):
        """
        :type c: int
        :rtype: bool
        """
        from math import sqrt
        l,h = 0,sqrt(c)//1
        while l<=h:
        	sum=(l**2)+(h**2)
        	if sum == c:
        		return True
        	elif sum < c:
        		l+=1
        	elif sum > c:
        		h-=1
        return False
