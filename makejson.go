package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var personName string
	var adrs string

	fmt.Println("Enter Your Full name")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	personName = scanner.Text()

	fmt.Println("Enter the address ")

	scanner.Scan()
	adrs = scanner.Text()

	var userTest = map[string]string{"fname": personName, "address": adrs}

	userTestJSON, err := json.Marshal(userTest)

	if err != nil {
		fmt.Println("Terjadi Kesalahan! Hehe")
	}

	fmt.Println("Json Output: ", string(userTestJSON))

}
