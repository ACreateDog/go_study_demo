package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c , cancel := context.WithTimeout(context.Background() , 1 * time.Second)
	c = context.WithValue(c,"k","v")
	c = context.WithValue(c ,"k1","v1")

	go handle(c ,"do some thing")
	cancel()
	select {
	case <-c.Done():
		fmt.Println("c done",c.Err())
	}
	time.Sleep(2 * time.Second)
}
func handle(c context.Context , doString string ) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println(doString)
	fmt.Println(c.Value("k"))
	fmt.Println(c.Value("k1"))
	select {
	case <-c.Done():
		fmt.Println("c done handle",c.Err() == context.Canceled , c.Err() == context.DeadlineExceeded)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("do something done")
	}
}
func subHandle(c context.Context, doString string) {

}