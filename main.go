package main

import (
	"demo/encript"
	"encoding/base64"
	"fmt"
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

func main() {

	key := "lYMa4WkYHmql2dlW"
	data := "Wade@2431313"
	ret :=  encript.EcbEncrypt([]byte(data) , []byte(key))
	r :=  base64.StdEncoding.EncodeToString(ret)
	fmt.Println(r)



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
