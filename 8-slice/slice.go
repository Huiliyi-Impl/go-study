package main

import "fmt"

func main() {
	myArray := []int{1, 2, 3}
	myArray2 := []int{11, 22, 33}
	myArray3 := make([]int, 10)
	myArray = append(myArray, myArray2...)
	myArray3 = append(myArray3, myArray2...)
	PrintArray2(myArray)
	fmt.Println(myArray)
	fmt.Println(myArray3)
}

// PrintArray2 slice引用传递
func PrintArray2(myArray []int) {
	for index, value := range myArray {
		fmt.Println(value)
		myArray[index] = value * 2
	}
}
