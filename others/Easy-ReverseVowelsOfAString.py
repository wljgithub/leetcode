'''

Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:
Given s = "hello", return "holle".

Example 2:
Given s = "leetcode", return "leotcede".

Note:
The vowels does not include the letter "y".

'''
class Solution:
	def reverseVowels(self, s):
		"""
		:type s: str
		:rtype: str
		"""
		print(s)
		l,h = 0,len(s)-1        
		vowels=['a','e','i','o','u']
		while l<h:
			if s[l] in vowels and s[h] in vowels:
				print(s[l],s[h])
				# s[h] = 1
			while s[l] not in vowels:
				l+=1 
			while s[h] not in vowels:
				h-=1
		return s

o=Solution()
print(o.reverseVowels('leetcode'))
