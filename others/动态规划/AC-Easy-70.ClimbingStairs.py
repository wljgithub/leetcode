'''
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
'''

d = {}

class Solution(object):
	def climbStairs_1(self, n):
		"""
		time complexity: O(2^n)
		this algorithm exceed the 
		time 
		"""
		if n <= 2:
			return n
		else:
			return self.climbStairs_1(n-1) + self.climbStairs_1(n-2)

	def climbStairs_2(self, n):
		"""
		time complexity: O(n)
		space complexity: O(n)
		"""
		if n <=2 :
			return n
		
		elif n in d:
			return d[n]
		else:
			
			val = self.climbStairs_2(n-1)+self.climbStairs_2(n-2)
			d [n] = val
			return val
	def climbStairs_3(self,n):
		'''
		time complexity: O(n)
		space comlexity: O(1)
		'''
		if n < 3:
			return n 
		a,b = 1,2            
		for i in range(n-3):
			a,b = b,a+b 
		return a+b

o=Solution()        
print(o.climbStairs_3(2))