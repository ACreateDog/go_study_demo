package main

import (
	"fmt"
	"strconv"
)

type Ainterface interface {

	AddFunc()
}

type Imp struct {

}

func (imp Imp) AddFunc() {
	fmt.Println("ahahah")
}
func DoString(a Ainterface)  {
	if a == nil {
		fmt.Println("a is nil")
	}
	a.AddFunc()
}

func main() {
	var i Imp
	DoString(i)
	ret :=  strconv.FormatInt(1575964259766-200, 10)
	fmt.Println(ret)

}