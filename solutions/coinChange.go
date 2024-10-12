package main

import (
	"fmt"
	// "string"
)
func coinChange(coins []int, amount int) int {
    res := make([]int, amount+1)

    for i := range res {
        res[i] = amount +1
    }
    res[0] = 0


    for i := 1; i < len(res); i++ {
        for coins_index := range coins {
            curr_coin := coins[coins_index]
            if curr_coin <= i {
                prevBestIndex :=  i - curr_coin
                prevBest := res[prevBestIndex]
                if prevBest < res[i] {
                    res[i] = prevBest +1
                }

            }
        }
    }
    fmt.Println(res)
    if res[amount] != amount + 1 {
        return res[amount]  

    } else {
        return -1
    }
}

func main() {
    // 3
    coins := []int{1,2,5}
    amount := 11
    x1 := coinChange(coins, amount)
    fmt.Println("result: ",x1)

    // 2
    coins = []int{1}
    amount = 2
    x1 = coinChange(coins, amount)
    fmt.Println("result: ",x1)

    // 0
    coins = []int{1}
    amount = 0
    x1 = coinChange(coins, amount)
    fmt.Println("result: ",x1)

    // -1
    coins = []int{2}
    amount = 3
    x1 = coinChange(coins, amount)
    fmt.Println("result: ",x1)
}


	

// 1, 2, 6, 24
	
	