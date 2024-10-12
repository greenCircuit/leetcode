package main

import (
	"fmt"
	// "string"
)

func twoSum(nums []int, target int) []int {
    count := make(map[int]int)
    fin := make([]int, 2)
    // building map
    for i :=0; i < len(nums); i++{
        currnum := nums[i]
        need := target - currnum
        val, ok := count[need]
        // If the key exists
        if ok {
            fin[0] = i
            fin[1] = val 
            return fin
        } else {
            count[currnum] = i
        }
    }
    return fin
}

func main() {
    arr1 := []int{1, 2, 3, 4}
    x1 :=twoSum(arr1, 4)
    fmt.Println(x1)
    fmt.Println("=====================")
    arr2 := []int{2,7,11,15}
    x2 :=twoSum(arr2, 9)
    fmt.Println(x2)
    fmt.Println("=====================")
    arr3 := []int{3, 2, 4}
    x3 :=twoSum(arr3, 6)
    fmt.Println(x3)
    fmt.Println("=====================")
    arr4 := []int{3, 3}
    x4 :=twoSum(arr4, 6)
    fmt.Println(x4)
    fmt.Println("=====================")
}


	

// 1, 2, 6, 24
	
	