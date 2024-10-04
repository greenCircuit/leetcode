package main

import (
	"fmt"
	// "string"
)

func main() {
    count := 5
    main_arr := make([][]int, count)
    init_arr := []int{1}
    main_arr[0] =  init_arr

    main_arr[1] =  []int{1, 1}
    // fmt.Println(main_arr)
    for i := 2; i < count; i++ {
        fmt.Println("count: ", i)
        new_arr := make([]int, i+ 1)
        new_arr[0] = 1
        new_arr[i] = 1
        fmt.Println("new arr", new_arr)
        prev_array := main_arr[i-1]
        fmt.Println("prev arr", prev_array)


        fmt.Println(new_arr)
        for j :=1; j < i; j++ {
            fmt.Println(prev_array)
            new_number := prev_array[j] + prev_array[j-1]
            fmt.Println("new number:", new_number)
            new_arr[j] = new_number
            
        }
        // fmt.Println(new_arr)
        main_arr[i] = new_arr
            // fmt.Println(prev_array)
    // fmt.Println("=======================")
    }
    fmt.Println(main_arr)
}


	


	