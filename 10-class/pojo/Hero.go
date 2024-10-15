package pojo

import "fmt"

type Hero struct {
	Name   string
	Age    int
	Skill  string
	Weapon string
}

func (h *Hero) GetName() string {
	return h.Name
}
func (h *Hero) SetName(name string) {
	h.Name = name
}
func (h *Hero) walk() {
	fmt.Println("walk")
}
func (h *Hero) run() {
	fmt.Println("run")
}
