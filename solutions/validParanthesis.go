package main

import (
	"fmt"
	// "string"
)

func isValid(s string) bool {
    if (len(s)) % 2 !=0 {
        return false
    }

    var stack []byte

	for i := range s {
        fmt.Println(stack)

		curr_char := s[i]
		if curr_char == '(' || curr_char == '[' || curr_char == '{' {
			stack = append(stack, curr_char)
		} else {
			if len(stack) == 0 {
				return false
			} else {

				curr_size := len(stack)
                if curr_char == '}' && stack[curr_size -1] == '{' {
				    stack = stack[:curr_size-1]
                } else if curr_char == ')' && stack[curr_size -1] == '(' {
				    stack = stack[:curr_size -1]
                } else if curr_char == ']' && stack[curr_size- 1] == '[' {
				    stack = stack[:curr_size -1]
                } else {
                    return false
                }
			}
		}

	}
	return len(stack) == 0
}

func main() {
	x := "()"
	x1 := isValid(x)
	fmt.Println("result: ", x1)

	x = "[}"
	x1 = isValid(x)
	fmt.Println("result: ", x1)

	x = "([])"
	x1 = isValid(x)
	fmt.Println("result: ", x1)
    
	x = "(("
	x1 = isValid(x)
	fmt.Println("result: ", x1)
}

// 1, 2, 6, 24
