package main

import (
	"container/list"
	"time"
	"sync"
	"fmt"
)
type waitingItem struct{
	RobotID string //机器人ID
	Ttl  int64     //超时时间
}

var wg sync.WaitGroup
func main() {
	l := list.New()


	ele := l.PushBack(2)
	//l.PushBack(2)
	//l.PushBack(2)

	//ele := l.Front()
	//tch := time.Tick(5 * time.Second)
	//monitor(&tch)


	//ticker := time.NewTicker(5 * time.Second)
	//monitor(&ticker.C)
	//closeCh(ticker)
	//ticker.Stop()
	//wg.Wait()
	l.Remove(ele)
}

func monitor(ch * <-chan time.Time) {

	wg.Add(1)
	go func( ch * <-chan time.Time) {
		defer wg.Done()
		d:= <-(*ch)
		fmt.Println("时间到 hello world : " , d)
	}(ch)
}


//func closeCh(ticker * time.Ticker) {
//	time.Sleep(2 * time.Second)
//	ticker.Stop()
//}
