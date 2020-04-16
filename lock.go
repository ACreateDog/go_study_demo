package main

import (
	"demo/common"
	"sync"
	"time"
	"fmt"
)

var rwlock sync.RWMutex
var mutex sync.Mutex

var incInt  int
var wg sync.WaitGroup

type demoS struct {
	Poo string

}

func (*demoS) dd() {
	fmt.Println("demo")
}

type subDemoS struct {
	demoS
}

func (*subDemoS) dd() {
	fmt.Println("subDemo")
}

type i interface {
	dd()
}

func do(d i)  {
	d.dd()
}

func main() {
	common.AddListener("a", func(info common.EventInfo) {
		fmt.Println(info)
	})

}
func main1()  {
	var a  = demoS{
		Poo: "a",
	}
	var d  = subDemoS{
		a,
	}
	fmt.Println(d.Poo)

	do(&a)


	go func() {
		rwlock.Lock()
		time.Sleep(15 * time.Second)
		defer rwlock.Unlock()
	}()
	go func() {
		time.Sleep(3 * time.Second)
		rwlock.RLock()
		fmt.Println("1 fetch rlock")
		defer rwlock.RUnlock()
	}()

	go func() {
		time.Sleep(4 * time.Second)
		rwlock.Lock()
		fmt.Println("1-fetch lock")
		defer rwlock.Unlock()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		rwlock.RLock()
		fmt.Println("1 fetch rlock")
		defer rwlock.RUnlock()
	}()
	time.Sleep(20 * time.Second)
}

func doString() {

	go func() {
		rwlock.RLock()
		time.Sleep(15 * time.Second)
		defer rwlock.RUnlock()
	}()

	go func() {
		time.Sleep(4 * time.Second)
		rwlock.RLock()
		fmt.Println("1 fetch rlock")
		defer rwlock.RUnlock()

	}()
	go func() {
		time.Sleep(3 * time.Second)
		rwlock.Lock()
		fmt.Println("1-fetch lock")
		defer rwlock.Unlock()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		rwlock.Lock()
		fmt.Println("2-fetch lock")
		defer rwlock.Unlock()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		rwlock.RLock()
		fmt.Println("2 fetch rlock")
		defer rwlock.RUnlock()

	}()
	go func() {
		time.Sleep(7 * time.Second)
		rwlock.Lock()
		fmt.Println("3-fetch lock")
		defer rwlock.Unlock()
	}()

	go func() {
		time.Sleep(8 * time.Second)
		rwlock.RLock()
		fmt.Println("3 fetch rlock")
		defer rwlock.RUnlock()
	}()
	go func() {
		time.Sleep(9 * time.Second)
		rwlock.Lock()
		fmt.Println("4-fetch lock")
		defer rwlock.Unlock()
	}()

	go func() {
		time.Sleep(10 * time.Second)
		rwlock.RLock()
		fmt.Println("4 fetch rlock")
		defer rwlock.RUnlock()
	}()



	time.Sleep(20 * time.Second)
	//t1 := time.Now().UnixNano()
	//fn(40)
	//t2 := time.Now().UnixNano()
	//fmt.Println((t2 - t1) / 1000000)

}
func fn( i int ) int {
	if i == 1 || i == 2  {
		return  1
	}
	return  fn(i-1) + fn(i - 2)
}

func selectTimes() {


}
func setLock() {
	//var lock bool
	//lock = true
	//
	//
	//fmt.Println(lock)


}


//
//func main() {
//	ch := make(chan []int, 6)
//	var m []int
//	for i := 0; i < 5; i++ {
//		go func() {
//			ch <- m
//			return
//		}()
//	}
//	for c := range ch {
//		fmt.Printf("c is %v \n", c)
//	}
//}

//func main() {
//	ch := make(chan []int, 4)
//	var m []int
//
//	var wg sync.WaitGroup
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func() {
//			ch <- m
//			return
//		}()
//	}
//	go func() {
//		for c := range ch {
//			fmt.Printf("c is %v\n", c)
//			wg.Done()
//		}
//	}()
//	wg.Wait()
//}



func wlk()  {
	mutex.Lock()
	incInt++
	fmt.Println("write some thing" , incInt)
	time.Sleep(10 * time.Second)
	mutex.Unlock()

	wg.Done()
}
func rlk()  {
	mutex.Lock()
	fmt.Println(" read some thing " , incInt)
	time.Sleep(5 * time.Second)
	mutex.Unlock()
	wg.Done()
}

func read() {

	for n := 0; n < 5 ; n++ {
		time.Sleep(1 * time.Second)
		rwlock.Lock()
		//incInt++
		fmt.Println(9)
		rwlock.Unlock()
	}

	wg.Done()
}
func rlock() {

	rwlock.RLock()
	fmt.Println("rlock....." , incInt)
	time.Sleep(time.Second * 10)

	defer rwlock.RUnlock()
	wg.Done()
}