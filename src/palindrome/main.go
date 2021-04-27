package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

//	Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

//	Note: For the purpose of this problem, we define empty string as valid palindrome.

//Example 1:
//		Input: "A man, a plan, a canal: Panama"
//		Output: true

//Example 2:
//		Input: "race a car"
//		Output: false

func main() {
	fmt.Println(isPalindrome("anNA"))
	fmt.Println(isPalindrome("a man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Printf("%s\n", strings.Repeat("#", 50))
	fmt.Println(isPalindromeSanitize("anNA"))
	fmt.Println(isPalindromeSanitize("a man, a plan, a canal: Panama"))
	fmt.Println(isPalindromeSanitize("race a car"))
}

func isPalindrome(str string) bool {
	chars := []rune(str)
	l := len(chars)
	for start, end := 0, l-1; start < end ; start, end = start+1, end-1 {
		for start < end && !unicode.IsLetter(chars[start]) {
			start++
		}
		for start < end && !unicode.IsLetter(chars[end]) {
			end--
		}
		if strings.ToLower(string(chars[start])) != strings.ToLower(string(chars[end])) {
			return false
		}
	}
	return true
}

func isPalindromeSanitize(str string) bool {
	result := sanitize(str)
	l := len(result)
	for i := 0; i < l-1; i++ {
		if result[i] != result[l-i-1] {
			return false
		}
	}
	return true
}

func sanitize(str string) string {
	reg, _ := regexp.Compile(`[^A-Za-z0-9]`)
	result := reg.ReplaceAllString(str, "")
	return strings.Trim(strings.ToLower(result), "")
}
