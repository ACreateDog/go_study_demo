package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

type consumer struct {

}
func main ()  {
	Consumer()
}
func Consumer() {
	cfg := nsq.NewConfig()

	c , err :=  nsq.NewConsumer("test" , "chan1" , cfg)
	if err != nil {
		fmt.Printf("error=%+v", err)
		return
	}

	c.AddHandler(nsq.HandlerFunc(
		func(message * nsq.Message) error {
			fmt.Printf("message=%+v",message)
			return nil
		}))

	time.Sleep(1000 * time.Second)
}

//func (c consumer) HandleMessage( message * nsq.Message )  error {
//	fmt.Printf("message=%+v",message)
//	return nil
//}
