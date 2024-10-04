package main

import (
	"fmt"
	// "string"
)

func hasDuplicate(nums []int) bool {
    tmp := make(map[int]int)

	for _, v := range nums {
        
		tmp[v]++
        
		if tmp[v] > 1 {
			return true
		}
	}
    fmt.Println(tmp)


    return false
}

func main() {
    arr1 := []int{1, 2, 3, 4, 5}
    x1 := hasDuplicate(arr1)
    fmt.Println(x1)

    arr2 := []int{1, 2, 3, 4, 5}
    x2 := hasDuplicate(arr2)
    fmt.Println(x2)
}


	


	
	