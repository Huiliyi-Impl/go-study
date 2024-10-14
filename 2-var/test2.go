package main

import "fmt"

func main() {
	var a int = 1
	a++
	fmt.Println(a)
	fmt.Printf("type of a = %T\n", a)

	var b = 100
	fmt.Println(b)
	fmt.Printf("type of b = %T\n", b)

	var c string = "go"
	fmt.Printf("value of c = %s ,type of c = %T\n", c, c)
	//2-var d string = "language"

	e := "100.00"
	fmt.Printf(e)

	var xx, yy int = 100, 200
	fmt.Println(xx, yy)

	var (
		aaa int    = 100
		bbb string = "hello"
	)
	fmt.Println(aaa, bbb)
}
