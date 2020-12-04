package main

import (
	"encoding/json"
	"fmt"
	"github.com/json-iterator/go"
	"sync"
	"time"
)

type In interface {
	DoAction1()
	DoAction2()
	DoAction3()
}

type AP struct {
	Foo string
	Db  sync.Map
}

func (a *AP) Get(k string) (interface{}, bool) {
	return a.Db.Load(k)
}
func (a *AP) Set(k string, v interface{}) {
	a.Db.Store(k, v)
}
func (AP) DoAction1() {
	fmt.Println("AP do action1")
}

func (AP) DoAction2() {
	fmt.Println("AP")
}

func (AP) DoAction3() {
	fmt.Println("AP")
}

type A struct {
	AP
	Foo  string
	Foo1 int
	PP   []string
}

func (a A) GetPP() []string {
	return a.PP
}
func (a *A) DoAction1() {
	a.AP.DoAction1()
	fmt.Println("A DoAction1")
}

type TmpFoo struct {
	AP
}

type Robot struct {
	// 锁
	sync.Mutex

	RobotID      string `json:"robotID"`
	HouseID      string `json:"warehouseID"`
	RegionID     string `json:"regionID"`
	PauseStatus  int    `json:"pauseStatus"` //机器人暂停状态，0：正常；1：暂停
	OnlineState  string `json:"onlineState"`
	TesState     string `json:"tesState"`
	AgvState     string `json:"agvState"`
	CurrentPoint string `json:"currentPoint"`
	Ext          string `json:"ext"`
	WhoLock      string `json:"whoLock"`
	RobotType    string `json:"robotType"`
	BelongTo     int8   `json:"belongTo"`
	DeviceSN     string `json:"deviceSN"`
	DeviceType   string `json:"deviceType"`
	//通用
	PodTheta   float32  `json:"podTheta"`
	Theta      float32  `json:"theta"`
	CurX       int32    `json:"curX"`    //当前x坐标 agvext
	CurY       int32    `json:"curY"`    //当前y坐标
	CurZ       int32    `json:"curZ"`    // 当前z坐标
	CurDir     uint8    `json:"curDir"`  //当前朝向
	MapID      uint32   `json:"mapID"`   //当前Z坐标
	PodID      string   `json:"frameID"` //机器人顶举的货架号
	LiftState  uint8    `json:"liftState"`
	Sequence   uint64   `json:"seq"`          //状态序列，用于保存最新状态
	SideADir   uint8    `json:"sideADir"`     //货架A面方向
	Timestamp  uint64   `json:"timestamp"`    //时间戳，用于前端显示
	ErrorState []string `json:"errorState"`   //具体异常原因，每个字符串代表一个异常
	ErrorMsg   []string `json:"errorMessage"` //具体异常原因，每个字符串代表一个异常
	UcPower    uint8    `json:"ucPower"`      //剩余电量, %
	AutoToRest int      `json:"autoToRest"`   // 是否自动回休息区
	HaveCar    int      `json:"haveCar"`      // 提升机专用，提升机内是否有机器人
	ManualLock string   `json:"manualLock"`   //'停车作业锁定状态'
	LockStatus string   `json:"lockStatus"`
	UcAlarm    uint32   `json:"ucAlarm"`
	LiftTheta  float64  `json:"liftTheta"`
	LiftHeight int      `json:"liftHeight"`
}

func main() {

	jsonstr := `{
    "robotID": "123",
    "warehouseID": "327238215972945955",
    "regionID": "325301802956226573",
    "pauseStatus": 0,
    "onlineState": "ONLINE",
    "tesState": "OK",
    "agvState": "IDLE",
    "currentPoint": "158286130280",
    "ext": "{\"curDir\":0,\"curX\":6100,\"curY\":18450,\"errorMessage\":null,\"errorState\":[],\"frameId\":65535,\"liftHeight\":0,\"liftState\":0,\"liftTheta\":0,\"mapID\":2001,\"mapVer\":\"1.0.0.0\",\"payloadKg\":0,\"rotateModel\":0,\"seq\":1604580497849,\"sideADir\":1,\"status\":0,\"temper\":0,\"theta\":1.5707964,\"timestamp\":1604581596269,\"ucAlarm\":0,\"ucPower\":100,\"ucState\":0,\"ucStateSave\":0,\"ucStatus\":0,\"ucWork\":0,\"ulRunTime\":0,\"uniqueID\":0,\"usSpeed\":0,\"voltage\":0}",
    "whoLock": "TES",
    "robotType": "0",
    "belongTo": 0,
    "deviceSN": "123",
    "deviceType": "0",
    "podTheta": 0,
    "theta": 1.5707964,
    "curX": 6100,
    "curY": 18450,
    "curZ": 0,
    "curDir": 0,
    "mapID": 2001,
    "frameID": "65535",
    "liftState": 0,
    "seq": 1604580497849,
    "sideADir": 1,
    "timestamp": 1604581596269,
    "errorState": [],
    "errorMessage": null,
    "ucPower": 100,
    "autoToRest": 0,
    "haveCar": 0,
    "manualLock": "",
    "lockStatus": "",
    "ucAlarm": 0,
    "liftTheta": 0,
    "liftHeight": 0
}`
	t1 := time.Now().UnixNano()
	for i := 0; i < 100000; i++ {
		var r Robot
		json.Unmarshal([]byte(jsonstr), &r)
	}
	t2 := time.Now().UnixNano()

	fmt.Println("json,cost=", (t2-t1)/1e6, "ms")

	t1 = time.Now().UnixNano()
	for i := 0; i < 100000; i++ {
		var r Robot
		e := jsoniter.Unmarshal([]byte(jsonstr), &r)
		if e != nil {
			fmt.Println(e)
			break
		}
	}
	t2 = time.Now().UnixNano()

	fmt.Println("jsoniter,cost=", (t2-t1)/1e6, "ms")

	// jsoniter 3687068000
	// json      649938000
	return

	var aa = []int{1, 2, 3, 4, 5, 6, 7}
	var b = aa[2:6]
	fmt.Println(len(b), b)

	var foo interface{}
	foo = AP{
		Foo: "foo",
		Db:  sync.Map{},
	}

	tmp := foo.(AP)
	fmt.Println(tmp)

	return
	// 16009334331-5-1-3-4
	var fooo []string
	var a A
	//a.PP = []string{"a","b"}
	fooo = append(fooo, a.GetPP()...)
	fmt.Println(fooo)
	return
	//foo := new(A)
	//foo.Foo = "foo"
	//foo.DoAction1()

}
