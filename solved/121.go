// Say you have an array for which the ith element is the price of a given stock on day i.
// 
// If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
// 
// Note that you cannot sell a stock before you buy one.
// 
// Example 1:
// 
// Input: [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
//              Not 7-1 = 6, as selling price needs to be larger than buying price.
// Example 2:
// 
// Input: [7,6,4,3,1]
// Output: 0
// explanation: In this case, no transaction is done, i.e. max profit = 0.

// 暴力破解 时间复杂度O(n^2)
func maxProfit(prices []int) int {
    max:=0
    for i,v:=range prices{
        for j:=i;j<len(prices);j++{
            if temp:=prices[j]-v;temp>max{
                max = temp
            }
        }
    }
    return max
}

// onepass O(n)
func maxProfit(prices []int) int {
    minPrice:=getMax(prices)
    maxProfit:=0

    for _,v:=range prices{
        if v<minPrice{
            minPrice = v
        }else if (v-minPrice)>maxProfit{
            maxProfit = v - minPrice
        }
    }
    return maxProfit
}

func getMax(arr []int)(max int){
    if len(arr)<1{
       return 0
    }

    max = arr[0]
    for _,v:=range arr[1:]{
        if v>max{
            max = v
        }
    }
    return max
}
