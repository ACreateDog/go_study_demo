package common

import (
	"fmt"
	"runtime"
)

var listenerMap = make(map[string][]func( EventInfo))

const EventCommandSuccess = "CommandSuccess"
const EventCommandFail = "CommandFail"
const EventRegisterOnline = "RegisterOnline" //机器人注册事件
const EventOffline = "Offline"               //机器人下线
const EventRobotOnChangeTypePower = "RobotOnChangePower"     //机器人在某个点
const EventRobotOnChangeTypePoint  = "RobotOnChangePoint"
const EventRobotOnline = "RobotOnline"
const CommandComplete  = "CommandComplete"
const EventRobtDataReport  = "EventRobtDataReport"

type EventInfo struct {
	Data  interface{}
	EventType string
}

func AddListener(EventType string, callback func( EventInfo)) {
	a , b , c ,d := runtime.Caller(1)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	listenerList := listenerMap[EventType]
	if len(listenerList) <= 0 {
		listenerMap[EventType] = []func( EventInfo){}
	}
	//每个事件类型是一个列表，可以容纳多个事件
	listenerMap[EventType] = append(listenerMap[EventType], callback)

}

func CallListener(eventType string, dat interface{}) {
	params := EventInfo{
		Data : dat ,
		EventType : eventType ,
	}

	callbackFuncList := listenerMap[eventType]
	if len(callbackFuncList) <= 0 {
		return
	}
	//循环调度
	for _ , callbackFunc := range callbackFuncList {
		callbackFunc( params)
	}
}

