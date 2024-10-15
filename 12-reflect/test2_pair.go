package main

import "fmt"

type Reader interface {
	ReadBook()
}
type Writer interface {
	WriteBook()
}
type Book struct {
}

func (b *Book) ReadBook() {
	fmt.Println("ReadBook")
}
func (b *Book) WriteBook() {
	fmt.Println("WriteBook")
}
func main() {
	b := &Book{}
	var r Reader
	r = b
	r.ReadBook()

	var w Writer
	w = b
	w.WriteBook()
}
