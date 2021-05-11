package main

import (
	"expvar"
	"fmt"
	"net/http"
	"sync"
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
type TestMap  map[string]int
type Foo struct {
	TestMap
	lock sync.RWMutex
	F *InF
	N string
}
type Error struct {
	ErrorCode    int
	ErrorMsg     string
	ErrorUserMsg string
}

func NewError(code int, message string, userMessage string) Error {
	return Error{
		ErrorCode:    code,
		ErrorMsg:     message,
		ErrorUserMsg: userMessage,
	}
}

func (err Error) Error() string {
	return err.ErrorMsg
}

func (err Error) Foo() string {
	return "foo"
}
var  errorTopicNoClient = Error{ErrorCode: 5000001, ErrorMsg: "this topic no match fd", ErrorUserMsg: "没有对应的会话通道"}

func GetErr()error  {
	return errorTopicNoClient
}

type TmpFoo struct {
	Val string
}

type WFError struct {
	Error
}
func main() {
	var e WFError

	fmt.Println(e.Error.Error())
	e.Foo()



	var tmpMap = map[string]TmpFoo{
		"a":TmpFoo{"a"},
		"b":TmpFoo{"b"},
		"c":TmpFoo{"c"},
		"d":TmpFoo{"d"},
	}
	var ret  = make(map[string]*TmpFoo)
	for _,t := range tmpMap {
		ret[t.Val] = &t
	}
	fmt.Println(ret)

	//connErr := make(chan int)
	//for {
	//	select {
	//	case <-connErr:
	//
	//	default:
	//	}
	//	fmt.Println(1000)
	//}

	//var  x = 0
	//go func() {
	//	for i := 0 ; i < 10000 ; i ++ {
	//		x = 1
	//		fmt.Println("100",x)
	//
	//	}
	//}()
	//go func() {
	//	x = 0
	//}()
	//
	//time.Sleep(2*time.Second)

	//var tmpMap = make(map[string][]string)
	//a := tmpMap["aa"]
	//fmt.Println(len(a))
	//
	//
	//return
	//arr1 := [3]int{1, 2, 3}
	//arr2 := [...]int{1, 2, 3}
	//fmt.Println(arr1 == arr2)
	//f := new(Foo)
	//f.N = "aaa"
	//f.TestMap= make(map[string]int)
	//fmt.Println(*f)

	//var sss = []string{"a","b","c"}
	//for _,n := range sss{
	//	go func(s string) {
	//		fmt.Println(s)
	//	}(n)
	//}
	//time.Sleep(5 * time.Second)

	//var a  = &Foo{}
	//fmt.Printf("%+v\n",a)
	//var d sync.Map
	//fmt.Println(d.LoadOrStore("a","vv"))
	//fmt.Println(d.LoadOrStore("a","vvvv"))
	if GetErr() == errorTopicNoClient {
		fmt.Println(errorTopicNoClient)
	}
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
