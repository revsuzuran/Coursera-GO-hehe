package main

import "fmt"

func main() {
	var x float64
	fmt.Println("Enter float number :")
	fmt.Scan(&x)
	fmt.Println("The truncated = ", int64(x))
}
