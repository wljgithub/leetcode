class Solution:

    
    def titleToNumber(self, s):
        """
        :type s: str
        :rtype: int
        """
        sum=0
        for i in range(len(s)):
            sum+=self.turnToNum(s[-(i+1)])*(26**i)
        return sum

    def turnToNum(self,s):
        import string
        d={string.ascii_uppercase[i]:i+1 for i in range(len(string.ascii_lowercase))}
        return d[s] 

o=Solution()
# o.turnToNum('A')
print(o.titleToNumber('AFDSG'))