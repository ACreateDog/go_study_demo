package main

import (
	"sync"
	"fmt"
	"os"
)

type Instance struct {
	foo string
}

var singleLock sync.Mutex
var singleOnce sync.Once
var wg sync.WaitGroup
func main() {
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			 ins := GetInstance2()
			 fmt.Printf("%p\n" , ins)
			 adrr := fmt.Sprintf("%p" , ins)

			 wg.Done()
		}()

	}
	wg.Wait()
}
var singleInstance *Instance

func GetInstance1() *Instance {
	if singleInstance == nil {
		singleLock.Lock()
		defer singleLock.Unlock()

		if singleInstance == nil {
			singleInstance = &Instance{
				foo:"hhhh" ,
			}
		}
	}
	return singleInstance
}

func GetInstance2() *Instance {
	singleOnce.Do(func() {
		singleInstance = &Instance{
			foo:"hhhh" ,
		}
	})
	return  singleInstance
}

func writeToFile(data string) {

	f , _ :=  os.Create("~/Desktop/tmp")
	f.Sync()
	f.WriteString(data)

}

