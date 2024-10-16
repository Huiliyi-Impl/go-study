package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := range 5 {
			c <- i
		}

		close(c)
	}()

	for data := range c {
		fmt.Println("data = ", data)
	}

	//for {
	//	// 判断是否关闭
	//	if num, ok := <-c; ok {
	//		fmt.Println("num = ", num)
	//	} else {
	//		break
	//	}
	//}
}
