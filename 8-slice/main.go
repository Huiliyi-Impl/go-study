package main

import (
	"fmt"
)

var a int = 11
var b int = 22
var c int = 33

func main() {
	var myArray [10]int
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [...]int{}
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
	for i := 0; i < len(myArray2); i++ {
		fmt.Println(myArray2[i])
	}
	for index, value := range myArray2 {
		fmt.Println(index, value)
	}
	fmt.Println(myArray3)
	re1, re2 := printArray(myArray2)
	fmt.Println(re1, re2)
}

func printArray(myArray [10]int) (r1 int, r2 int) {
	var sum int
	for value := range myArray {
		sum += value
	}
	return sum, len(myArray) - 1
}
