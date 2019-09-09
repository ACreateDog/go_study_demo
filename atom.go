package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

var cnt int64
var waitg sync.WaitGroup

func main() {
	i1 := 2.5 * math.Pi * 1000
	i2 := 2 * math.Pi * 1000
	int1 := int(i1)
	int2 := int(i2)

	fmt.Println(int1 % int2)
	fmt.Println(float64(int1 % int2) / 1000)

	//var a int64 = 22
	//ret := float64(a) * 1.5
	//fmt.Println(ret)

	//for i := 0 ; i < 10 ;i++ {
	//	waitg.Add(1)
	//	go addForAtomic()
	//}
	//
	//waitg.Wait()
	//fmt.Println(cnt)
}

func addForAtomic() {

	defer waitg.Done()
	atomic.AddInt64(&cnt , 1)
}
func syncCond() {
	var mutex sync.Mutex

	sync.NewCond(&mutex)

}
