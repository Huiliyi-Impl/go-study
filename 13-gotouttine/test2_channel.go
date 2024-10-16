package main

import (
	"fmt"
	"time"
)

func main() {
	//定义要给channel
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine 结束>>>")
		fmt.Println("goroutine 正在运行>>>")
		time.Sleep(5 * time.Second)
		c <- 100
	}()
	num := <-c //从c中读取数据
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束>>>")
}
