package main

import (
	"fmt"
	"github.com/wangfeiso/rwlock"
	"math/rand"
	"sync"
	"time"
)

func main() {

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
	i := 100
	var wg sync.WaitGroup
	wg.Add(200)
	for {
		i--
		go func(i int , wg *sync.WaitGroup) {
			lock := rwlock.New("uuuuuuu")
			lock.Lock()
			fmt.Println("Lock--",i)
			time.Sleep(1 * time.Second)
			lock.Unlock()
			fmt.Println("Unlock----",i)
			defer wg.Done()
		}(i , &wg)
		if i == 0 {
			break
		}
	}
	go func() {
		for {
			lock := rwlock.New("uuuuuuu")
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			lock.RLock()
			fmt.Println("RLock--")
			lock.RUnlock()
			fmt.Println("RUnlock----")
		}

	}()
	wg.Wait()
	//lock.RLock()
	//fmt.Println("lock")
	//time.Sleep(30 * time.Second)
	//lock.RUnlock()
	//fmt.Println("unlock")
}
