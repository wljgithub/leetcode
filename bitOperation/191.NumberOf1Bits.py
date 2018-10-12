'''
编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。

示例 :

输入: 11
输出: 3
解释: 整数 11 的二进制表示为 00000000000000000000000000001011
 

示例 2:

输入: 128
输出: 1
解释: 整数 128 的二进制表示为 00000000000000000000000010000000
'''
class Solution(object):
    # 这种写法如果输入的是负数，就会陷入死循环
    def hammingWeight1(self, n):
        """
        :type n: int
        :rtype: int
        """
        count =0
        while n>=1:
            if n&1==1:
                count+=1
            n = n>>1
            
        return count
    def hammingWeight2(self, n):
        """
        :type n: int
        :rtype: int
        """
        count =0
        flag =1
        while flag:
            if n&1==1:
                count+=1
            n = n>>1
            
        return count
    def hammingWeight3(self, n):
        """
        :type n: int
        :rtype: int
        """
        count =0
        while n!=0:
            count+=1
            n = n&(n-1)
            
        return count

if __name__ == '__main__':
    o=Solution()
    print(o.hammingWeight1(0x80000000))