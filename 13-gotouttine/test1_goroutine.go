package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		time.Sleep(time.Second)
		fmt.Printf("new Goroutine : i = %d\n", i)
		i++
	}
}
func main() {
	go newTask()

	//i := 0
	//
	//for {
	//	i++
	//	fmt.Printf("main goroutine : i = %d\n", i)
	//	time.Sleep(time.Second)
	//}
}
