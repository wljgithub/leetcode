'''
颠倒给定的 32 位无符号整数的二进制位。

示例:

输入: 43261596
输出: 964176192
解释: 43261596 的二进制表示形式为 00000010100101000001111010011100 ，
     返回 964176192，其二进制表示形式为 00111001011110000010100101000000 。
'''
class Solution:
    # @param n, an integer
    # @return an integer
    def reverseBits(self, n):
    	ret=0
    	for i in range(31):
    		ret=ret|(n&1)
    		ret<<=1
    		n>>=1
    	return ret

if __name__ == '__main__':
	o=Solution()
	print(o.reverseBits(0x00000010100101000001111010011100))