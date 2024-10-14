package main

import (
	"fmt"
	"time"
)

func main() {
	Test1()
	Test2()
	Test1()
}
func Test1() {
	fmt.Println("hello world")
}
func goFunc(i int) {
	fmt.Println("goroutine", i, ">>>")
}
func Test2() {
	for i := range 10000 {
		go goFunc(i)
	}
	time.Sleep(time.Second)
}
