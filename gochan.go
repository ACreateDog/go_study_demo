package main

import (
	"fmt"
	"sync"
	"time"
)

var glChan = make(chan int , 0)
var wg sync.WaitGroup

func main() {

	var a  map[string]string

	go func() {

		for n := 0; n < 10; n++ {
			glChan <- n
		}
	}()
	for i := 0; i < 10 ; i++ {
		go goSelectChanDothing(i)
	}

	time.Sleep(2 * time.Second)
}

func goSelectChanDothing(tag int) {

	for {
		select {
		case glInt := <-glChan:
			fmt.Println( tag , "--",glInt)

		}
	}
}

