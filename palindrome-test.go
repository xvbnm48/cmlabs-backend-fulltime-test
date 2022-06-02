package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var kata string
	fmt.Print("Masukkan kata: ")
	// fmt.Scan(&kata)
	b := bufio.NewScanner(os.Stdin)
	b.Scan()
	result := palindrome(b.Text())
	if result == true {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}
func palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
