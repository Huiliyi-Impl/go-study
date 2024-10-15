package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg is string type, value:", value)
	} else {
		fmt.Println("arg is not string type")
		fmt.Printf("value type is %T\n", value)
	}
	fmt.Println("=======================")
}

type Book struct {
	auth string
}

func main() {
	var GoBook = Book{"Go语言"}
	JavaBool := Book{"Java语言"}
	myFunc(GoBook)
	myFunc(JavaBool)
	myFunc("Go语言")
	myFunc(123)
	myFunc(true)
}
