package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

var (
	//nsqd的地址，使用了tcp监听的端口
	tcpNsqdAddrr = "127.0.0.1:4150"
)

func main() {

	produceer()
}
func produceer() {
	config := nsq.NewConfig()
	p  , err  :=  nsq.NewProducer(tcpNsqdAddrr , config)
	if err != nil {
		log("NewProducer",err)
		return
	}
	err = p.Publish("testTopic" , []byte("hello world"))
	if err != nil {
		log("Publish",err)
		return
	}
}
func log(data ...interface{}) {
	fmt.Printf("data=%+v",data)

}