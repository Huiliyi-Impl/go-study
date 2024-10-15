package main

import (
	"10-class/pojo"
	"fmt"
)

type Book struct {
	name   string
	author string
	price  float64
	pages  int
}

func main() {
	//var book Book
	//book.name = "Go语言"
	//book.author = "Go语言"
	//book.price = 100.00
	//book.pages = 100
	//fmt.Println(book)
	//changeBook(&book)
	//fmt.Println(book)
	//fmt.Println(changeBook2(book))
	//fmt.Println(book)
	hero := pojo.Hero{
		Name:   "hero",
		Age:    18,
		Skill:  "fire",
		Weapon: "gun",
	}
	hero.GetName()
	hero.SetName("hero2")
	fmt.Println(hero)

	superman := pojo.Superman{
		Hero:       hero,
		SuperPower: "fly",
	}
	superman.Walk()
	superman.Run()
}

func changeBook(book *Book) {
	book.name = "Java语言"
	book.author = "Java语言"
	book.price = 200.00
	book.pages = 200
}
func changeBook2(book Book) Book {
	book.name = "C语言"
	book.author = "C语言"
	book.price = 300.00
	book.pages = 300
	return book
}
