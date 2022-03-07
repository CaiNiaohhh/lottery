package main
func maxProfit(prices []int) int {
	Max, Min := 0, 9999
	for i := 0; i < len(prices); i++ {
		if Max < prices[i] - Min {
			Max = prices[i] - Min
		}
		if Min > prices[i] {
			Min = prices[i]
		}
	}
	return Max
}
