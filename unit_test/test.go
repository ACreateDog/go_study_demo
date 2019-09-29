package main

import "fmt"

func main() {
	ret := Fni(5)
	fmt.Println(ret)
}

func Fni( n  int) int   {
	if n < 2  {
		return  n
	}
	return Fni(n-1) + Fni(n-2)
}
