package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var swg sync.WaitGroup
var dataInt []int
var once sync.Once

//func fff() {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println(" print fff ")
//			panic("出错啦 ~ 抛出")
//		}
//	}()
//
//	panic("出错啦 ~ ")
//}

func main() {
	t := 39000000
	endT :=   time.Duration(t) / time.Millisecond
	fmt.Println(endT)
	os.Exit(0)
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(" print main  ")
	//		debug.PrintStack()
	//	}
	//}()
	//fff()
	//for i := 0; i < 10; i++ {
	//
	//	go func() {
	//		once.Do(onced)
	//		fmt.Println("213")
	//	}()
	//}
	//for  i , v := range make([]string , 10) {
	//	once.Do(onces)
	//	fmt.Println("count:", v, "---", i)
	//
	//
	//}
	//time.Sleep(4000)

	//i := big.NewInt(70)
	//rand.Int()


	//rabbitmq.Send()
	//rabbitmq.Receive()
}

func onces() {

	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
