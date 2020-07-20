package main

import (
	"fmt"
	"unsafe"
)

type Num struct {
	i string
	j int64
	k float64
}

func main()  {
	/*
		intNum := 5
		intNumPointer := &intNum

		floatNumPointer := (*float64)(unsafe.Pointer(intNumPointer))
		fmt.Println(*floatNumPointer)
	*/

	intArray()
}
func structPointer() {
	n := &Num{i:"DFASFASD",j:1,k:1.78}
	nPointer := unsafe.Pointer(n)
	niPointer := (*string)(unsafe.Pointer(nPointer))
	*niPointer = "卧槽~无情~"

	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(Num{}.j)))
	*njPointer = 2

	nkPointer := (*float64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(Num{}.k)))
	*nkPointer = 3.14

	fmt.Println(*n)
}
func intArray() {
	var intList = []int{1,2,3,4,5,6}
	intListPointer := unsafe.Pointer(&intList)
	item := (*int)(unsafe.Pointer(uintptr(intListPointer)+1))
	fmt.Println(*item)
}