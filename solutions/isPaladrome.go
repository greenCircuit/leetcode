package main

import (
	"fmt"
    "regexp"
	// "string"
)

func isAlpha(char string) bool {
    is_alphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(char)
    return is_alphanumeric
}

func isPalindrome(word string) bool {
    fmt.Println("finding if palindrome for: ", word)
    if len(word) == 1 {
        return true
    }
    
    a_pointer := 0
    b_pointer := len(word) - 1

    fmt.Println(a_pointer)
    fmt.Println(b_pointer)
    for a_pointer < b_pointer {
        a_char := word[a_pointer]
        b_char := word[b_pointer]

        // find next valid a_pointer char
        for a_pointer < b_pointer && !isAlpha(string(a_char)) {
            a_pointer += 1
        }
        
        // find next valid b_pointer char
        for b_pointer > a_pointer && !isAlpha(string(b_char)) {
            b_pointer -= 1
        }

        if word[a_pointer] == word[b_pointer] {
            a_pointer += 1
            b_pointer -= 1
        } else {
            return false
        }
    }

    return true
}

func main() {
    test_str := "aacb-caa"
    x1 :=isPalindrome(test_str)
    fmt.Println(x1)
}


	


	