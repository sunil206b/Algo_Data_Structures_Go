package main

import (
	"fmt"
	"strings"
)

func main() {
	reversedStr := reverseString("traditional for loop")
	fmt.Printf("String reversed with traditional for loop: %s\n", reversedStr)
	fmt.Println(strings.Repeat("=", 80))
	reversedStr = reverseStringWithRange("for range loop")
	fmt.Printf("String reversed with range loop: %s\n", reversedStr)
	fmt.Println(strings.Repeat("=", 80))
	reversedStr = reverseStringWithLength("reverse with length")
	fmt.Printf("String reversed with max length: %s\n", reversedStr)
	fmt.Println(strings.Repeat("=", 80))
	reversedWords := reverseWords("This is reversed by traditional for loop")
	fmt.Printf("Reversed words with traditional loop: %s\n", reversedWords)
	fmt.Println(strings.Repeat("=", 80))
	reversedWords = reverseWordsWithRange("This is reversed by range loop")
	fmt.Printf("Words reversed by for range loop: %s\n", reversedWords)
	fmt.Println(strings.Repeat("=", 80))
	reversedWords = reverseWordsWithLength("This is reversed  by using length in reverse")
	fmt.Printf("Words reversed with max length: %s\n", reversedWords)
	fmt.Println(strings.Repeat("=", 80))
	reversedWords = reverseWordsWithoutAPI("This is reversed without API methods")
	fmt.Printf("This is reversed without API: %s\n", reversedWords)
}

func reverseString(str string) string {
	chars := []rune(str)
	length := len(chars)
	if length < 1 {
		return ""
	}
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func reverseStringWithRange(str string) string {
	reversed := ""
	for _, s := range str {
		reversed = string(s) + reversed
	}
	return reversed
}

func reverseStringWithLength(str string) string {
	chars := []rune(str)
	length := len(chars)
	reversed := ""
	if length < 1 {
		return ""
	}
	for i := length - 1; i >= 0; i-- {
		reversed += string(chars[i])
	}
	return reversed
}

func reverseWords(str string) string {
	words := strings.Fields(str)
	length := len(words)
	if length < 1 {
		return ""
	}
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func reverseWordsWithLength(str string) string {
	words := strings.Fields(str)
	length := len(words)
	reversed := ""
	if length < 1 {
		return reversed
	}
	for i := length - 1; i >= 0; i-- {
		reversed += words[i] + " "
	}
	return strings.TrimSpace(reversed)
}

func reverseWordsWithRange(str string) string {
	words := strings.Fields(str)
	reversed := ""
	for _, w := range words {
		reversed = w + " " + reversed
	}
	return strings.TrimSpace(reversed)
}

func reverseWordsWithoutAPI(str string) string {
	chars := []rune(str)
	result := ""
	length := len(chars)
	i := 0
	for i < length {
		for i < length && chars[i] == ' ' {
			i++
		}
		if i >= length {
			break
		}
		j := i + 1
		for j < length && chars[j] != ' ' {
			j++
		}
		sub := chars[i:j]
		if len(result) == 0 {
			result = string(sub)
		} else {
			result = string(sub) + " " + result
		}
		i = j + 1
	}
	return result
}
