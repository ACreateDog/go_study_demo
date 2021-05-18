package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

type Observer interface {
	Notify(string)
}
type BaseOb struct {
	Ob * Observed
}
type OberserA struct {
	BaseOb
}

func (oa * OberserA) Notify( val string) {
	fmt.Println("OberserA ",val)
}
type OberserB struct {
	BaseOb
}

func (oa * OberserB) Notify(val string) {
	fmt.Println("OberserB ", val)
}



type Observed struct {
	Foo string
	ObserveList []Observer
}

func (o * Observed) SetFoo(val string) {
	o.Foo = val
	o.NotifyAll(val)
}
func (o *Observed) NotifyAll(val string) {
	for _,ob := range o.ObserveList {
		ob.Notify(val)
	}
}
func (o *Observed) AddNotifys(observer ...Observer) {
	o.ObserveList = append(o.ObserveList,observer...)
}



func main() {
	var observed Observed
	observed.AddNotifys(new(OberserA),new(OberserB))
	observed.SetFoo("hello wanghuanyi")


	//observableFun()
	//observableForEach()
	//time.Sleep(1000 * time.Second)
}
func observableFun() {
	observable := rxgo.Just("Hi,wangfei")()
	ch := observable.Observe()
	item := <-ch
	fmt.Println(item.V)
}
func observableForEach() {
	observable := rxgo.Just("Hi,wangfei")()
	for  i:=0 ; i < 10 ; i++   {
		observable.ForEach(func(v interface{}) {
			fmt.Printf("received: %v\n", v)
		}, func(err error) {
			fmt.Printf("error: %e\n", err)
		}, func() {
			fmt.Println("observable is closed")
		})

	}

	go func() {
		for i:=0 ; i < 10 ; i++ {
			rxgo.Just("Hi,wf")
			rxgo.Just("Hi,wf01")
			rxgo.Just("Hi,wf02")
			rxgo.Just("Hi,wf03")
		}
	}()
}
func producer() {
	//go rxgo.Producer()
}
