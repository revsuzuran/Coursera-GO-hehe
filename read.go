/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Enter filename: ")
	var filename string
	fmt.Scanln(&filename)

	ReadFile(filename)
}
func ReadFile(filename string) {

	fmt.Printf("\n\nReading a file in Go lang\n")

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", filename)
	fmt.Printf("\nData: %s", data)
	slice := make([]Name, 0, 3)
	var namee Name
	slice = append(slice, namee)
	fmt.Println("\n********************\n")

	for _, v := range slice {
		fmt.Println(v.fname, v.lname)
	}

}
