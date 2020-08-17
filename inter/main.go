package main

import "fmt"

type In interface {
	DoAction1()
	DoAction2()
	DoAction3()
}

type AP struct {
	Foo string
}

func (AP) DoAction1() {
	fmt.Println("AP")
}

func (AP) DoAction2() {
	fmt.Println("AP")
}

func (AP) DoAction3() {
	fmt.Println("AP")
}

type A struct {
	AP
	Foo string
	Foo1 int
}

func (a * A) DoAction1() {
	fmt.Println("A")
}






func main() {
	foo := new(A)
	foo.Foo = "foo"
	foo.DoAction1()
}
