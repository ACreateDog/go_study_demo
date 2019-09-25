package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main()  {
	str := `
今天是个好日子.
	但是下雨了。
把庄家淋湿了.
	哈哈哈
`
	fmt.Println(str)
	foo()
}
func foo() {
	str := "    hello  d   "

	str =  strings.Trim(str , " ")
	fmt.Println(str)
	f()
}

type User struct {
	Name string `json:"name"`
}
func f()  {
	var u User
	err :=  json.Unmarshal([]byte{} , &u)

	if err != nil {
		fmt.Printf("error=%+v",err)
	}
}