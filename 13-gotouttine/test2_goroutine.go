package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//用go创建继承一个形参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A-DEFER")

		func() {
			defer fmt.Println("B-DEFER")
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
		time.Sleep(time.Second)
	}
}
