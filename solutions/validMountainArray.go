package main

import (
	"fmt"
	// "string"
)

func validMountainArray(arr []int) bool {
	if len(arr) <= 2 {
		return false
	}
	peaks := 0
	for c := 1; c < len(arr) -1; c++ {
		fmt.Println(arr[c])
		fmt.Println(arr[c+1])
		
		if arr[c] == arr[c-1] || arr[c] == arr[c+1]{
			return false
		}

		if arr[c] < arr[c-1] && arr[c] < arr[c+1]{
			fmt.Println("valley found",arr[c])
			return false
		}

		if arr[c] > arr[c-1] && arr[c] > arr[c+1]{
			peaks++
		}
	}
	fmt.Println("peaks", peaks)
	return peaks == 1
}

func main() {
    coins := []int{1,2,5, 3}
	x1 := validMountainArray(coins)
	fmt.Println(x1)

	// coins = []int{0,1,2,1,2}
	// x1 = validMountainArray(coins)
	// fmt.Println(x1)
	
	coins = []int{0,3,2,1}
	x1 = validMountainArray(coins)
	fmt.Println(x1)

	// coins = []int{0,3}
	// x1 = validMountainArray(coins)
	// fmt.Println(x1)

	// coins = []int{3,5,5}
	// x1 = validMountainArray(coins)
	// fmt.Println(x1)
}


