package main

import (
	"fmt"
)

//Given an integer x, return true if x is palindrome integer.
//
//An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.
//
//Example 1:
//	Input: x = 121
//	Output: true
//Example 2:
//	Input: x = -121
//	Output: false
//Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
//Example 3:
//	Input: x = 10
//	Output: false
//Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

//Example 4:
//	Input: x = -101
//	Output: false

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(-101))
}

func isPalindrome(number int) bool {
	// If the number is negative or if the last digit is 0
	// and first digit is not 0 then it is not a palindrome
	// (only possible if the number is 0)
	if number < 0 || (number%10 == 0 && number != 0) {
		return false
	}

	revertedNumber := 0
	for number > revertedNumber {
		revertedNumber = revertedNumber * 10 + number % 10
		number = number / 10
	}
	return number == revertedNumber || number == revertedNumber/10
}
