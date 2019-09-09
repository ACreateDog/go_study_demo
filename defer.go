package main

import (
	"fmt"
	"github.com/siddontang/go-log/log"
	"math"
	"strconv"
)

func main() {
	//demo1()
	//fffff()
	fmt.Println(fmt.Sprintf("%f",math.Pi))
}

func demo1() {
	//var foo = make(map[string]interface{})
	//var listL []string
	//
	//foo["hhh"] = []string{}
	//foo["list"] = listL
	//
	//
	//fmt.Println(foo)
	//
	//Fo()
	t()
}
func Fo() int {
	var foo  = 8

	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()

	defer func() {
		fmt.Println("3")
	}()

	return  foo
}
func t() {
	r := 1029 & ^0XF

	fmt.Println(r)
}
func fffff() {
	a , err :=  strconv.Atoi("(x+y)/v")
	if err != nil {
		log.Printf("err is %+v" , err)
	}
	fmt.Println(a)
}