package main

import "fmt"

type Person interface {
	Walk()
	Speak()
}

type Batman struct {
	Name string
}

func (b Batman) Walk() {
	fmt.Println("I'm walking")
}

func (b Batman) Speak() {
	fmt.Println("I'm batman")
}

type Superman struct {
	Name string
}

func (s Superman) Walk() {
	fmt.Println("I'm walking")
}

func (s Superman) Speak() {
	fmt.Println("I'm superman")
}

func doSomething(p Person) {
	p.Walk()
	p.Speak()
}

func main() {
	b := Batman{"Batman"}
	s := Superman{"Superman"}

	doSomething(b)
	doSomething(s)
}
