package do1

import (
	"demo/common"
	"demo/do2"
	"fmt"
)

func init() {
	fmt.Println("init")
	common.AddListener(common.EventRobtDataReport , DoSomething)
}
func DoSomething(info common.EventInfo) {
	fmt.Println(info)
}
func DoSome() {
	//var p *int

	do2.DD()
}
