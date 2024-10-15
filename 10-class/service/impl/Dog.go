package impl

import "fmt"

type Dog struct {
	Color string
}

func (d *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}
func (d *Dog) GetColor() string {
	return d.Color
}
func (d *Dog) GetType() string {
	return "This is Dog"
}
