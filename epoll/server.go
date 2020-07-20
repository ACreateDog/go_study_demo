package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func base64Encode() {
	var  data = []byte("hello world")

	// output buff
	var buff  = make([]byte , base64.StdEncoding.EncodedLen(len(data)))
	// aGVsbG8gd29ybGQ=
	base64.StdEncoding.Encode(buff , data)
	fmt.Println(string(buff))

	// output string
	var ret = ""
	ret =  base64.StdEncoding.EncodeToString(data)
	fmt.Println(ret)

	//
	var a rune = 0xf0
	ret =  base64.StdEncoding.WithPadding(a).EncodeToString(data)
	fmt.Println(ret)

}

func lock() {
	var l = &sync.RWMutex{}
	//l.Unlock()
	//return
	//l.RLock()
	l.Lock()
	//l.RUnlock()
	//var wg = &sync.WaitGroup{}
	//wg.Add(1)
	//
	//l.Lock()
	//
	//go func(wg *sync.WaitGroup) {
	//	defer wg.Done()
	//	fmt.Println("lock ")
	//	l.Lock()
	//}(wg)
	//
	//wg.Wait()

}
func main() {
	//base64Encode()
	lock()
	time.Sleep(60 * time.Second)
	fmt.Println("exit")
	return
	StrConv()
	fmt.Println("hello string fmt")
	StrFmt()
}
func StrConv() {
	var intOne  =  103747892749090909
	t1 := time.Now().UnixNano()
	for i := 0 ; i < 1000000 ; i ++ {
		strconv.Itoa(intOne)
	}
	t := time.Now().UnixNano()
	fmt.Println(t - t1)
}

func StrFmt() {
	intOne := 103747892749090909
	t1 := time.Now().UnixNano()
	for i := 0 ; i < 1000000 ; i ++ {
		fmt.Sprintf("%+v",intOne)
	}
	t := time.Now().UnixNano()
	fmt.Println(t - t1)
}