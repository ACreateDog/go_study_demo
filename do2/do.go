package do2

import (
	"demo/common"
	"fmt"
)

func CallFunc() {
	fmt.Println("callFunc")
	common.CallListener(common.EventRobtDataReport , "hello")
	//do1.DoSomething(common.EventInfo{})
}
func DD()  {
	fmt.Println("funcName")
}