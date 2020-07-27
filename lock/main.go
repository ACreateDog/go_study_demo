package main

import (
	"fmt"
	"github.com/wangfeiso/rwlock"
	"math/rand"
	"sync"
	"time"
)

type Tmp struct {
	m sync.RWMutex
}

func main() {
	return
	var ch = make(chan string, 1)
	fmt.Println(len(ch))
	ch <- "aa"
	ch <- "dd"
	ch <- "cc"
	close(ch)
	go func() {
		for  {
			select {
			case dat := <-ch:
				fmt.Println(dat,"11111")
				for {
					select {
					case d := <- ch:
						fmt.Println(d)
					}
					continue
				}
				continue
			}
		}
		fmt.Println("closed chan")
	}()


	time.Sleep(10 * time.Second)
	//a := Tmp{}
	//a.m.RLock()
	//defer a.m.RUnlock()
	return
	//var l sync.RWMutex
	//
	//l.Lock()
	//go func(l *sync.RWMutex) {
	//	l.RLock()
	//	fmt.Println("rlock")
	//	defer l.RUnlock()
	//
	//}(&l)
	//time.Sleep(50 * time.Second)
	//l.Unlock()
	//time.Sleep(2 * time.Second)
	//return
	rwlock.Init(&rwlock.Options{
		Addr: "127.0.0.1:6379",
	})
	var lockCount int64= 0
	lock := rwlock.New("rqwerweqrqwrwq")
	t1 := time.Now().UnixNano()
	lock.Lock()
	t2 := time.Now().UnixNano()
	lockCount += t2 - t1

	fmt.Println("Lock--")
	time.Sleep(120 * time.Second)
	lock.Unlock()
	fmt.Println("Unlock----")
	//i := 1
	//var wg sync.WaitGroup
	//wg.Add(i)
	//for {
	//	i--
	//	go func(i int , wg *sync.WaitGroup) {
	//		lock := rwlock.New("afsfs9084924098")
	//		t1 := time.Now().UnixNano()
	//		lock.Lock()
	//		t2 := time.Now().UnixNano()
	//		lockCount += t2 - t1
	//
	//		fmt.Println("Lock--",i)
	//		time.Sleep(60 * time.Second)
	//		lock.Unlock()
	//		fmt.Println("Unlock----",i)
	//		defer wg.Done()
	//	}(i , &wg)
	//	if i == 0 {
	//		break
	//	}
	//}

	//go func() {
	//	for {
	//		lock := rwlock.New("uuuuuuu")
	//		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	//		lock.RLock()
	//		fmt.Println("RLock--")
	//		lock.RUnlock()
	//		fmt.Println("RUnlock----")
	//	}
	//
	//}()
	//wg.Wait()
	fmt.Println(lockCount)
	//lock.RLock()
	//fmt.Println("lock")
	//time.Sleep(30 * time.Second)
	//lock.RUnlock()
	//fmt.Println("unlock")
}
func Rand() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
}