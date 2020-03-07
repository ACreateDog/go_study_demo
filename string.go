package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func main()  {
	str := `
今天是个好日子.
	但是下雨了。
把庄家淋湿了.
	哈哈哈
`
	fmt.Println(str)

	type errorParam struct {
		RobotID   string
		ErrorCode int
		ErrorMsg  string
	}

	var msg = errorParam{}

	msg.RobotID = "6351"
	msg.ErrorCode = 110
	msg.ErrorMsg = "this is an error"

	strParam, _ := json.Marshal(&msg)
	fmt.Println(string(strParam))
	//foo()
	fmt.Println(time.Now().Unix())
}
func foo() {
	str := "    hello  d   "

	str =  strings.Trim(str , " ")
	fmt.Println(string(str))
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