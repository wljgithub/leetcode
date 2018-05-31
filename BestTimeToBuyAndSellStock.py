class Solution:
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        mini,mini_index=self.find_value_and_index(prices)
        if mini_index==len(prices)-1:
            return 0
        thelist=[i-mini for i in prices]
        print(thelist,mini)
        return max(thelist)
    def find_value_and_index(self,seq):
        minimum=min(seq)
        return minimum,seq.index(minimum)

if __name__ == '__main__':
    s=Solution()
    l=[7,1,5,3,6,4]
    print(s.maxProfit(l))