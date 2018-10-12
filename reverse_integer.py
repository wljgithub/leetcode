'''
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output:  321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
'''

class Solution:
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        
        x=str(x)
       
        #负数
        if len(x)==1:
        	return int(x)
        	
        if x[0]=='-':
            return int('-'+x[1:][::-1])

        #尾部含有0
        elif x[-1]=='0':
            return self.delete_zero(x)

        else:
        	return int(str(x)[::-1])

    def delete_zero(self,s):
        for i in range(1,len(s)+1):
            if s[-i]!='0':
                return int(s[:-i+1][::-1])


sl=Solution()
print(sl.reverse(0))
print(sl.reverse(101))
print(sl.reverse(100))
print(sl.reverse(-101))
print(sl.reverse(-100))
print(int("-001"))