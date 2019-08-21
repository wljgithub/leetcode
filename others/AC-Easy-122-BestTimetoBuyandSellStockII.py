class Solution:
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """							
        total_prires,day = 0,1
        while day < len(prices):
        	if prices[day] > prices[day-1]:
        		total_prires+=prices[day]-prices[day-1]
        	# print(total_prires)
        	day+=1
        return total_prires

o=Solution()
print(o.maxProfit([7,1,5,3,6,4]))