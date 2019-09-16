package main

import (
	"fmt"
	"unsafe"
)

type Poo struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	Address struct {
		Province string `json:"province"`
		Number string `json:"number"`
	}
}

func main() {
	//var f Poo
	//var s = "{}"
	//json.Unmarshal([]byte(s) ,&f)
	//if f == (Poo{}) {
	//	fmt.Print("this is  equel")
	//} else {
	//	fmt.Print("this is nots equel")
	//}
	var u1 U1
	var u2 U2
	var u3 U3
	var u4 U4
	var u5 U5
	fmt.Println("u1 size is ", unsafe.Sizeof(u1))
	fmt.Println("u1 size is ", unsafe.Sizeof(u2))
	fmt.Println("u1 size is ", unsafe.Sizeof(u3))
	fmt.Println("u1 size is ", unsafe.Sizeof(u4))
	fmt.Println("u1 size is ", unsafe.Sizeof(u5))

}
//内存对齐
type U1 struct {
	a  byte
	b  int64
	c  int32
}
type U2 struct {
	c  int32
	b  int64
	a  byte
}
type U3 struct {
	a  byte
	c  int32
	b  int64
}
type U4 struct {
	b  int64
	c  int32
	a  byte
}
//
type U5 struct {
	a  byte
	b  int32
	d  int16
	c  int32
}