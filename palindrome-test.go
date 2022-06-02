package main

import "fmt"

func main() {
	var kata string
	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)
	result := palindrome(kata)
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
