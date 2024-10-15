package pojo

import "fmt"

type Superman struct {
	Hero
	SuperPower string
}

func (s Superman) Walk() {
	fmt.Println("superman walk")
}

func (s Superman) Run() {
	fmt.Println("superman run")

}
