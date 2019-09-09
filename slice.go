package main

import (
  _ "net/http/pprof"
	"fmt"
)

func main() {
	//sli := []int{ 2 , 2 , 4 }
	s1 := make([]byte , 100)
	s2 := make([]byte , 10)
	s1 = append(s1 , s2...)
	fmt.Printf("%v" , s1)
	fmt.Printf("%v" , len(s1))

	//go func() {
	//
	//}()
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//
	//}()
	//log.Println(http.ListenAndServe("localhost:6068", nil))

}

