package main

import "fmt"

func main() {
	defer fmt.Println("main end")
	fmt.Println("main start")
	defer fmt.Println("defer1")
	defer func1()
	defer func2()
}
func func1() {
	fmt.Println("A")
}
func func2() {
	fmt.Println("B")
}
