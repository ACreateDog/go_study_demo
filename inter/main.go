package main

import (
	"fmt"
	"sync"
)

type In interface {
	DoAction1()
	DoAction2()
	DoAction3()
}

type AP struct {
	Foo string
	Db sync.Map
}

func (a *AP) Get(k string) (interface{} , bool)  {
	return a.Db.Load(k)
}
func (a *AP) Set(k string , v interface{})  {
	a.Db.Store(k, v)
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

type TmpFoo struct {
	AP
}




func main() {
	//foo := new(A)
	//foo.Foo = "foo"
	//foo.DoAction1()
	a := TmpFoo{}
	a.AP.Set("aaa","bbb")
	fmt.Println(a.Get("aaa"))

}
