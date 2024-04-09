package demo

import "fmt"

type Methods interface {
	A()
	B()
}

type Human struct {
	Head string
}

func (Human) A() {
	fmt.Println("A")
}
func (Human) B() {
	fmt.Println("B")
}

type Man struct {
	Beard string
	Human
}

func UseMan() {
	man := Man{
		Beard: "black",
		Human: Human{Head: "white"},
	}
	man.A()
}
