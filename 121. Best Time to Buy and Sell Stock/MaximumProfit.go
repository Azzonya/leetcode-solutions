package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var profit, minValue int = 0, prices[0]

	for i := 1; i < len(prices); i++ {
		//fmt.Println(minValue, prices[i])
		if minValue > prices[i] {
			minValue = prices[i]
		} else {
			if profit != 0 {
				//fmt.Println(profit)
				if profit < prices[i]-minValue {
					profit = prices[i] - minValue
				}
			} else {
				profit = prices[i] - minValue
			}
		}

	}

	return profit
}
