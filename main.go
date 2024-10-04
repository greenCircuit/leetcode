package main

import (
	"fmt"
	// "string"
)

func productofArrayExceptSelf(nums []int) []int {
    N := len(nums)
    // newBack := make([]int, len(arr))
    left_products := make([]int, N)
    right_products := make([]int, N)
    fin := make([]int, N)

    left_products[0] = 1
    right_products[N-1] = 1
    for i := 1; i < len(nums); i++{
        newNum := nums[i-1] * left_products[i-1]
        left_products[i] = newNum
    }

    for i := N - 2; i >= 0; i--{
        newNum := nums[i+1] * right_products[i+1]
        right_products[i] = newNum
    }
    // fmt.Println(left_products)
    // fmt.Println(right_products)

    for i :=0; i <N; i++{
        fin[i] = left_products[i] * right_products[i]
    }
    // fmt.Println(fin)
    return fin
}

func main() {
    arr1 := []int{1, 2, 3, 4}
    x1 :=productofArrayExceptSelf(arr1)
    fmt.Println(x1)
    fmt.Println("=====================")
    // arr2 := []int{1, 2, 3, 4, 5}
    // productofArrayExceptSelf(arr2)
}


	

// 1, 2, 6, 24
	
	