package main

import "fmt"

func main() {
	myArray := make([]int, 3)
	fmt.Printf("len = %d,cap = %d,array = %v\n", len(myArray), cap(myArray), myArray)

	myArray = append(myArray, 1)
	fmt.Printf("len = %d,cap = %d,array = %v\n", len(myArray), cap(myArray), myArray)

	myArray = append(myArray, 2, 3)
	fmt.Printf("len = %d,cap = %d,array = %v\n", len(myArray), cap(myArray), myArray)

	s1 := myArray[4:]
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(myArray)

	s3 := make([]int, len(myArray))
	copy(s3, myArray)
	fmt.Println(s3)
}
