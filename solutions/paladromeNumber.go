package main

import (
	"fmt"
	// "string"
)

func palindromeNum(numCheck int) bool{
    if numCheck < 0 {
        return false
    }
    reversed := 0
    numCheckOg := numCheck
    for numCheck > 0 {
        remainder := numCheck % 10
        reversed = reversed * 10 + remainder
        numCheck = numCheck /10 
    }
    return reversed == numCheckOg
}
func main() {
    x1 := palindromeNum(101)
    x2 := palindromeNum(11)
    fmt.Println(x1)
    fmt.Println(x2)
}


	


	