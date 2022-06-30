package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var test_num string
	arrSlice := make([]int, 3)
	for true {
		fmt.Printf("\nEnter Number (Enter 'X' to stop):")
		fmt.Scan(&test_num)
		if test_num == "X" || test_num == "x" {
			break
		}
		i, err := strconv.Atoi(test_num)
		if err == nil {
		}
		arrSlice = append(arrSlice, i)
		sort.Sort(sort.IntSlice(arrSlice))
		fmt.Printf(fmt.Sprint(arrSlice))
	}

}
