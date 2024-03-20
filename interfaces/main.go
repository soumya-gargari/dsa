package main

import "fmt"

type Singer interface {
	Sing()
}

type Name struct {
	name string
}

type Author struct {
	authorName string
}

type book struct {
	bookName string
}

func (n *Name) Sing() {
	fmt.Println("singer name is:", n.name)
}

func (a *Author) Sing() {
	fmt.Println("author name is:", a.authorName)
}

func (b *book) Sing() {
	fmt.Println("book name is:", b.bookName)
}

func introduce(s Singer) {
	s.Sing()
}

func main() {
	i1 := &Name{name: "Soumya Gargari"}
	i2 := &Author{authorName: "Ipsita Biswas"}
	i3 := &book{bookName: "Love story"}
	introduce(i1)
	introduce(i2)
	introduce(i3)
}
