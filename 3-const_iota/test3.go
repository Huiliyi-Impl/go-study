package main

import "fmt"

// const来对应枚举
const (
	//iota从0开始, 每行递增1
	HONGKONG = iota + 1
	BEIJING
	SHANGHAI = iota * 2
	CHENGDU
	HANGZHOU
	SICHUAN
)

func main() {
	const length int = 10
	const width int = 5
	fmt.Println(length * width)
	fmt.Println(HONGKONG)

}
