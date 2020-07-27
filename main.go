package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
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

type AI interface {
	A()
	B()
}
type St struct {
	Foo string
}

func (s *St) A() {
	//panic("implement me")
	s.Foo = "hello world"
}

func (s St) B() {
	s.Foo = "aaa"
}
var parseDisplay struct {
	TagID                       int    `json:"tagID" validate:"min=1,max=8000"`
	ColorType                   int    `json:"colorType" validate:"min=0,max=7"`
	//ColorType                   string    `json:"colorType" validate:"iscolor"`
	Content                     string `json:"content" validate:"number|lte=0"`
	//Content                     string `json:"content" validate:"number"`
	ContentChange               int    `json:"contentChange" validate:"min=0,max=1"`
	DisplayType                 int    `json:"displayType" validate:"min=1,max=3"`
	AfterPressConfirmScreenType int    `json:"afterPressConfirmScreenType" validate:"min=1,max=3"`
}
func main() {
	var dat1 Data
	dat1.Ip = "127.0.0.1"
	var dat2 Data
	dat2.Ip = "127.0.0.1"

	fmt.Println(dat1 == dat2)



	//cfg := &validator.Config{
	//	TagName:      "validate",
	//}
	//jsonStr := `{"tagID":100,"colorType":6,"content":"","contentChange":1,"displayType":2,"afterPressConfirmScreenType":1}`
	//err :=   json.Unmarshal([]byte(jsonStr),&parseDisplay)
	//fmt.Println(err)
	//
	//err =  validator.New(cfg).Struct(parseDisplay)
	//if err == nil {
	//	return
	//}
	//tmpErr := err.(validator.ValidationErrors)
	//for k ,v := range tmpErr {
	//	fmt.Println(k)
	//	 fmt.Printf("error=%+v\n",v)
	//}
	//fmt.Println(err)
	return
	demoChan := make(chan int , 1)
	dat := 1
	demoChan <- 1
	select {
	case demoChan<-dat:
		fmt.Println("can input ")
	default:
		fmt.Println("cannot input chan " , dat)

	}


	return
	var thisOne = 1
	switch thisOne {
	case  1:

	case 2:
		fmt.Println("this is two")
	default:
		fmt.Println(" this is nothing")
	}

	return
	buf := make([]byte, 10000)
	file, err := os.Open("./data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//n, err := io2.ReadFull(file, buf)
	n ,err :=  file.Read(buf)
	if err != nil {
		fmt.Println(n, err.Error())
	} else {
		fmt.Println(n)
	}
	fmt.Println(string(buf))

	return
	var ffoo = make(map[string]*Data)
	
	fmt.Println( ffoo["a"].Email)
	//var  aaa byte
	//aaa = byte(1) + byte(0x30)
	//fmt.Println(aaa)
	return
	var ai AI
	var st St
	ai  = &st
	ai.B()
	fmt.Println(st.Foo)
	ai.A()
	fmt.Println(st.Foo)
	return
	aInt := 3
	fmt.Println(fmt.Sprintf("%03d",aInt))
	return
	var fo = []byte{'0','0','3'}
	fmt.Println(strconv.Atoi(string(fo)))
	return
	var foo AI

	foo = nil
	fmt.Println(foo == nil)
	return
	curTime := time.Now()
	endTime := curTime.Format("2006-01-02 15:04:05")
	startTime := curTime.AddDate(0 , 0 , -3).Format("2006-01-02 15:04:05")

	fmt.Println(endTime)
	fmt.Println(startTime)
	return


		var aslice = []string{"a","b","c","d"}
		fmt.Println(aslice[0:2])


		var aa interface{}
		 aa = ""
		 //fmt.Println(reflect.ValueOf(aa).IsNil())
		fmt.Println(aa == "")

		var a map[string]int
		if  _ , ok := a["a"] ; !ok {
			fmt.Println("not exist")
		}

		fmt.Println(len(a))

		f := []Data{ {
			Email:        "1-111",
			Password:     "1-222",
			Token:        "1-333",
			Authenticate: "1-444",
			Ip:           "1-555",
		},
		 {
			Email:        "2-111",
			Password:     "2-222",
			Token:        "2-333",
			Authenticate: "2-444",
			Ip:           "2-555",
		}}
		r, _ := json.Marshal(f)
		fmt.Println(string(r))

	//var foo  = map[string]Data{
	//	"1": {
	//		Email:        "1-111",
	//		Password:     "1-222",
	//		Token:        "1-333",
	//		Authenticate: "1-444",
	//		Ip:           "1-555",
	//	},
	//	"2": {
	//		Email:        "2-111",
	//		Password:     "2-222",
	//		Token:        "2-333",
	//		Authenticate: "2-444",
	//		Ip:           "2-555",
	//	},
	//}
	//
	//var tmp = make(map[string]*Data)
	//for i , v := range foo {
	//	t := v
	//	tmp[i] = &t
	//}
	//
	//fmt.Println(tmp)



	//resp := "{\"returnCode\":0,\"returnMsg\":\"succ\",\"returnUserMsg\":\"成功\",\"data\":{\"pid\":258248667432026933}}"
	//
	//var respObj struct {
	//	ReturnCode    int                    `json:"returnCode"`
	//	ReturnMsg     string                 `json:"returnMsg"`
	//	ReturnUserMsg string                 `json:"returnUserMsg"`
	//	Data          map[string]int64 `json:"data"`
	//}
	//json.Unmarshal([]byte(resp), &respObj)
	//fmt.Println(respObj)
	//fmt.Println(respObj.Data["pid"])
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
