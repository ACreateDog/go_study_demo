package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var swg sync.WaitGroup
var dataInt []int
var once sync.Once

//func fff() {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println(" print fff ")
//			panic("出错啦 ~ 抛出")
//		}
//	}()
//
//	panic("出错啦 ~ ")
//}

type Data struct{
	Email string `json:"userEmail"`
	Password string `json:"code"`
	Token string `json:"token"`
	Authenticate string `json:"authenticate"`
	Ip string `json:"ip"`
}

type UserCenterParam struct {
	FromSys   string `json:"fromSys"`
	TransId   string `json:"transId"`
	TransTime string `json:"TransTime"`
	Data      Data   `json:"data"`
}
type aInf interface {
	Write()
}
type aS struct {
	 Foo  string
}

func (a  aS) Write() {
	a.Foo = "wwwww"
	fmt.Println("Write")
}
var data = "hjfklasjflsf567890kjhgfghjkl;56789op[poiuy"
var countW = 1000000

//3951983000
//395
func seekWrite() {

	f , _ :=  os.OpenFile("./seekwrite.txt",os.O_RDWR|os.O_CREATE, 0600)
	curWritePos := 0
	for i := 0 ;i <= countW;i++{
		f.Seek(int64(curWritePos),0)
		f.WriteString(data)
		curWritePos += len(data)
	}
}
// 1 000 000
//3312376000
//331
func appendWrite() {
	f , _ := os.OpenFile("./append.txt", os.O_RDWR | os.O_CREATE |os.O_APPEND,0600)
	for i := 0 ;i <= countW;i++{
		f.WriteString(data)
	}
}

func main() {
	resp := "{\"returnCode\":0,\"returnMsg\":\"succ\",\"returnUserMsg\":\"成功\",\"data\":{\"pid\":258248667432026933}}"

	var respObj struct {
		ReturnCode    int                    `json:"returnCode"`
		ReturnMsg     string                 `json:"returnMsg"`
		ReturnUserMsg string                 `json:"returnUserMsg"`
		Data          map[string]int64 `json:"data"`
	}
	json.Unmarshal([]byte(resp), &respObj)
	fmt.Println(respObj)
	fmt.Println(respObj.Data["pid"])
	//dqLogf := func(level diskqueue.LogLevel, f string, args ...interface{}) {
	//	fmt.Printf(f,args)
	//}
	//// backend names, for uniqueness, automatically include the topic...
	//
	//backend := diskqueue.New(
	//	"backendName",
	//	"./data",
	//	10000000,
	//	1,
	//	100,
	//	5,
	//	10,
	//	dqLogf,
	//)
	//for i:=0; i <= 1000 ; i++  {
	//	err :=  backend.Put([]byte("hello world"))
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	time.Sleep(500 * time.Millisecond)
	//}

	//t1 := time.Now().UnixNano()
	//seekWrite()
	////appendWrite()
	//t2 := time.Now().UnixNano()
	//fmt.Println(t2 - t1)
	//fmt.Println((t2 - t1)/10e6)
	//var a interface{}
	//a = false
	//
	//if a.(bool) {
	//	fmt.Println(a)
	//}

	//pipline := make(chan string)
	//go func() {
	//	pipline <- "hello world"
	//	pipline <- "hello China"
	//	//close(pipline)
	//}()
	//for data := range pipline{
	//	fmt.Println(data)
	//}

	//var t aS
	//var tt  interface{} = t
	////var a = "string foo"
	////var b = 9
	////var c = 'a'
	//var a aS
	//a.Foo = "bbb"
	//fmt.Println(a.Foo)
	//a.Write()
	//fmt.Println(a.Foo)
	//switch tt.(type) {
	//case aInf:
	//	fmt.Println("as")
	//case string:
	//	fmt.Println("string")
	//default:
	//
	//	fmt.Println("default type")
	//}
	//println( reflect.TypeOf(a).String() )
	//println(reflect.TypeOf(b).String())
	//println(reflect.TypeOf(c).String())





	//key := "lYMa4WkYHmql2dlW"
	//data := "Wade@2431313"
	//ret :=  encript.EcbEncrypt([]byte(data) , []byte(key))
	//r :=  base64.StdEncoding.EncodeToString(ret)
	//fmt.Println(r)



	//tmp := UserCenterParam{
	//	FromSys:"",
	//	TransId:"",
	//	TransTime:"",
	//	Data:Data{
	//		Email:"",
	//		Password:"",
	//	},
	//}
	//t :=  UserCenterParam{
	//	FromSys:   "",
	//	TransId:   "",
	//	TransTime: "",
	//	Data:      Data{
	//		Email:        "",
	//		Password:     "",
	//		Token:        "",
	//		Authenticate: "",
	//		Ip:           "",
	//	},
	//}
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//fmt.Println(time.Now().Unix())


	//var m = make(map[string]int)
	//m["a"] = 1
	//b := m["b"]
	//fmt.Println(b)
	//
	//var pools []redsync.Pool
	//p :=  redis.Pool{
	//	Dial: func() (conn redis.Conn, e error) {
	//		conn , e =  redis.Dial("tcp" , "127.0.0.1:6379")
	//
	//		if e != nil {
	//			return nil, e
	//		}
	//		return conn , e
	//	},
	//	TestOnBorrow: func(c redis.Conn, t time.Time) error {
	//		_ , err := c.Do("PINT")
	//		return err
	//	},
	//	MaxIdle:         3,
	//	IdleTimeout:     240 * time.Millisecond,
	//}
	//
	//pools = append(pools  , &p)
	//
	//reds :=  redsync.New(pools)
	//redsync.SetTries(100)
	//redsync.SetRetryDelay(time.Duration(1000) * time.Millisecond)
	//mutex :=  reds.NewMutex("LockKey")
	//mutex.Lock()
	//mutex.Unlock()






	//var a interface{}
	//if a == "" {
	//	fmt.Println("true")
	//} else {
	//	fmt.Println("false")
	//}

	//do1.DoSome()
	//do2.CallFunc()
	//t := 39000000
	//endT :=   time.Duration(t) / time.Millisecond
	//fmt.Println(endT)
	//os.Exit(0)
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(" print main  ")
	//		debug.PrintStack()
	//	}
	//}()
	//fff()
	//for i := 0; i < 10; i++ {
	//
	//	go func() {
	//		once.Do(onced)
	//		fmt.Println("213")
	//	}()
	//}
	//for  i , v := range make([]string , 10) {
	//	once.Do(onces)
	//	fmt.Println("count:", v, "---", i)
	//
	//
	//}
	//time.Sleep(4000)

	//i := big.NewInt(70)
	//rand.Int()


	//rabbitmq.Send()
	//rabbitmq.Receive()
}

func onces() {

	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
