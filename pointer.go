package main

import "fmt"

type Person struct {
	Age int		`json:"age"`
	Name string `json:"name"`
}

func (p *  Person) GrowUp() {
	p.Age++
}

func main() {

	var person  = new(Person)
	person.Age = 10
	fmt.Println( "成长之前", person.Age)

	person.GrowUp()

	fmt.Println( "成长之后", person.Age)

}
