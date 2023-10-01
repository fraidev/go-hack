package main

import (
	"fmt"
	"strings"
)

func main() {
	result := reverseWords("Let's take LeetCode contest")
	fmt.Println(result)
}

func reverseWords(s string) string {
	var result string

	words := strings.Split(s, " ")

	for i, s := range words {
		reverseWord := reverseWord(s)

		result = result + reverseWord
		if i != len(words)-1 {
			result = result + " "
		}
	}

	return result
}

func reverseWord(s string) string {
	var word []byte
	lenS := len(s)
	for i := lenS; i > 0; i-- {
		c := s[i-1]
		word = append(word, c)
	}
	return string(word)
}
