package main

import (
	"fmt"
	"encoding/json"
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
	var f Poo
	var s = "{}"
	json.Unmarshal([]byte(s) ,&f)
	if f == (Poo{}) {
		fmt.Print("this is  equel")
	} else {
		fmt.Print("this is nots equel")
	}

}
