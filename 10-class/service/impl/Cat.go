package impl

import "fmt"

type Cat struct {
	Color string
}

func (c Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (c Cat) GetColor() string {
	return c.Color
}

func (c Cat) GetType() string {
	return "This is cat"
}
