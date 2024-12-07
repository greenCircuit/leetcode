package main

import (
	"fmt"
	// "string"
)

func maxArea(height []int) int {


	a_pointer := 0
	maxArea := 0
	b_pointer := len(height) - 1


	for a_pointer <= b_pointer {

		minHigh := height[a_pointer]
		if height[b_pointer] < height[a_pointer] {
			minHigh = height[b_pointer]
		}

        maxAreaTemp := (b_pointer - a_pointer) * minHigh
		
		if maxAreaTemp > maxArea {
			maxArea = maxAreaTemp
		}


		if height[a_pointer] < height[b_pointer] {
			a_pointer++
		} else {
			b_pointer--
		}
	}
	return maxArea
}



func main() {
	x1 := []int{1,8,6,2,5,4,8,3,7}
	x2 := maxArea(x1)
	fmt.Println(x2)

}


