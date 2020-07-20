package main

import (
	"fmt"
	"os"
	exec "os/exec"
	"strings"
	"time"
)
var data = `statistics,wid=303432559562326030,subWid=303432559562326030,deviceType=robotScroll timestamp=1592295531180945261,chargingCount=0,exceptionCount=0,idleCount=1,offlineCount=0,workingCount=0,pausedCount=0,errorInfo="{\"0\":[],\"256\":[],\"1\":[],\"517\":[],\"1029\":[],\"2\":[],\"3\":[],\"4\":[],\"6\":[],\"7\":[],\"8\":[],\"9\":[],\"768\":[],\"12\":[],\"1024\":[]}",task="{\"failedCount\":2}" 1592295531180945261
`
var  filePath = "/Users/wangfei/work/src/demo/data.txt"
func appendWrite() {


	f , err  := os.OpenFile(filePath, os.O_RDWR | os.O_CREATE |os.O_APPEND ,0600)
	fmt.Println(err)
	for i := 0 ;i <= 3000000;i++{
		f.WriteString(data)
	}
	fmt.Println(f.Sync())
}
func main() {


	t1 := time.Now().UnixNano()
	//cleanFile(filePath)
	t2 := time.Now().UnixNano()
	fmt.Println((t2 - t1) / 1e6)


	appendWrite()
}
func cleanFile(filePath string) {
	var ret []byte
	var err error
	cmd := exec.Command("/bin/sh", "-c", `echo "" > ` + filePath)
	if ret, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(ret))
	fmt.Println(strings.Trim(string(ret), "\n"))
}