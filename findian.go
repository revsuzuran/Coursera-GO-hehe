package main

import (
	"fmt"
	"strings"
)

func findian() {
	var uInput string
	fmt.Println("Enter Text :")
	fmt.Scan(&uInput)
	findWord(strings.ToLower(uInput))
}

func findWord(someWord string) {

	isStringStartWithI := strings.HasPrefix(someWord, "i")
	isStringEndWithN := strings.HasSuffix(someWord, "n")
	isStringContainsA := strings.Contains(someWord, "a")

	if isStringStartWithI && isStringEndWithN && isStringContainsA {
		fmt.Println("Found!", someWord)
	} else {
		fmt.Println("Not Found!", someWord)
	}

}
