package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) //带缓冲的channel

	go func() {
		defer fmt.Println("goroutine 结束>>>")

		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("goroutine 正在运行>>>,发送的元素", i, "len(c)=", len(c), ",cap(c)=", cap(c))
		}
	}()

	for i := 0; i < 10; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("main goroutine 结束>>>")
}
