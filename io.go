package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const FileName  = "/Users/wangfei/work/demo/tmp/tmp.txt"
const BaseDir  = "/Users/wangfei/work/demo/tmp/"
func main() {
	//osOpen()
	//ioUtil()
	//ioUtilWrite()
	//osWirte()
	osOpen()
}
func readBuf() {

}
func ioUtil()  {
	dat , err :=  ioutil.ReadFile(FileName)
	if err != nil {
		fmt.Printf("ReadFile is error %v" , err)
	}

	fmt.Printf("%v" , string(dat))
}
func ioUtilWrite() {
	by := []byte("this is block data")

	err := ioutil.WriteFile(BaseDir + "log1.txt" , by , 0644)
	if err != nil {
		fmt.Printf("error wirte %+v " , err)
	}
}
func osWirte() {

	fd  , err :=  os.Create(BaseDir + "tmp_write.txt")
	if err != nil {
		fmt.Printf("os.Create is error %v" , err)
	}
	n , err := fd.WriteAt([]byte("string string") , 5)
	if err != nil {
		fmt.Printf("fd.WriteAt is error %v" , err)
	}

	fmt.Printf("write data n %v" , n)

}
func osOpen() {
	fd , err :=  os.Open(FileName)

	fmt.Printf("%+v\n" , fd)
	if err != nil {z
		fmt.Printf("open is error is %v \n" , err)
		return
	}
	var buf  = make([]byte  , 2000)
	n , err :=  fd.Read(buf)
	if err != nil {
		fmt.Printf("Read is error is %v \n" , err)
		return
	}
	fmt.Printf("\n%+v\n" , string(buf[:n]))
	fmt.Println(n)
}
func ioCloser() {

}
func io() {

}
