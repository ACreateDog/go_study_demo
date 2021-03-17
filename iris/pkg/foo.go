package pkg

import "fmt"
var Foo = "this is foo"
func init() {
	fmt.Println("2")
}

func init() {
	fmt.Println("1")
}

