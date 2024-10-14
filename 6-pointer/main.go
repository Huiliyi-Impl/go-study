package main

import "fmt"

func main() {
	var a int = 5
	var b int = 6
	changeValue(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
func changeValue(a *int, b *int) {
	tmp := *a
	fmt.Printf("type of tmp is %T\n", tmp)
	*a = *b
	*b = tmp

}
