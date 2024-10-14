package main

import "fmt"

func main() {
	var re = foo1(10, 20)
	fmt.Println(re)
	var (
		result1, result2 = foo2("hello", "world")
	)
	ret1, ret2 := foo2("hello", "world")
	fmt.Println(result1, result2)
	fmt.Println(ret1, ret2)

	a, b := foo3("hello go!", 10022)
	fmt.Println(a, b)
}
func foo1(a int, b int) int {
	fmt.Println(a, b)
	var c int = a * b
	return c
}
func foo2(a string, b string) (string, string) {
	return a + b, "hello"
}
func foo3(a string, b int) (r1 string, r2 int) {
	r1 = a
	r2 = b
	return
}
