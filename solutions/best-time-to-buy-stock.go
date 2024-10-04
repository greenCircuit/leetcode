package main

import (
	"fmt"
	// "string"
)

func stock(prices []int) int {
    maxProfit := 0
    minPrice := prices[0]
    for i :=1; i < len(prices); i++ {
        fmt.Println("curr num: ", prices[i])
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else {

            tempProfit := prices[i] - minPrice
            if tempProfit > maxProfit {
                maxProfit = tempProfit
            }
        }
    }


    return maxProfit
}

func main() {
    x1 := []int{5,4, 5, 3, 11, 10}
    fmt.Println("max profit: ", stock(x1))

    x2 := []int{4, 3, 1, 1}
    fmt.Println("max profit: ", stock(x2))
}


	


	