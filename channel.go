package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)
	//c <- 1
	//c <- 2
	//c <- 3
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ClearActionWithChannel,Recovered=", r)
		}
	}()
	close(c)
	close(c)

	i, isClose := <-c
	if !isClose {
		fmt.Println("channel closed!")
		//break
	}
	fmt.Println(i)
	//for {
	//}
}