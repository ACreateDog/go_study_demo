package main

import (
	"expvar"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	s    map[string]string
	user User
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func showMap() interface{} {
	return s
}

func showUser() interface{} {
	return user
}

type InF interface {
	DoDo()
}
type Foo struct {
	lock sync.RWMutex
	F *InF
	N string
}
func main() {

	var sss = []string{"a","b","c"}
	for _,n := range sss{
		go func(s string) {
			fmt.Println(s)
		}(n)
	}
	time.Sleep(5 * time.Second)

	//var a  = &Foo{}
	//fmt.Printf("%+v\n",a)
	//var d sync.Map
	//fmt.Println(d.LoadOrStore("a","vv"))
	//fmt.Println(d.LoadOrStore("a","vvvv"))
	return
	user.Name = "tom"
	user.Age = 18
	s = make(map[string]string, 10)
	s["1"] = "111"
	s["2"] = "222"
	expvar.Publish("a_map", expvar.Func(showMap))
	expvar.Publish("b_user", expvar.Func(showUser))
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		s["3"] = "333"
	})
	http.ListenAndServe(":8088", nil)
}
