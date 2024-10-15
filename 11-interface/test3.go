package main

import "fmt"

type AnimaIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (c *Cat) Sleep() {
	fmt.Println("cat is sleep")
}
func (c *Cat) GetColor() string {
	return c.color
}
func (c *Cat) GetType() string {
	return "cat"
}

type Dog struct {
	color string
}

func (d *Dog) Sleep() {
	fmt.Println("dog is sleep")
}
func (d *Dog) GetColor() string {
	return d.color
}
func (d *Dog) GetType() string {
	return "dog"
}
func showAnimal(animal AnimaIF) {
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())
}
func main() {
	//var animal AnimaIF
	//animal = &Cat{"black"}
	//animal.Sleep()
	//animal = &Dog{"white"}
	//animal.Sleep()
	cat := Cat{"black"}
	dog := Dog{"white"}
	showAnimal(&cat)
	showAnimal(&dog)
	fmt.Println("============")
	var animal AnimaIF = &Cat{"black"}
	animal.Sleep()
	fmt.Println(animal.GetColor())
}
