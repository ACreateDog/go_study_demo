package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	//fhandle , _  := os.Open("")
	//fd := fhandle.Fd()
	var buf bytes.Buffer
	dataLen := 10
	if  err :=  binary.Write(&buf, binary.BigEndian, int64(dataLen)) ; err != nil {
		fmt.Println(err)
	}
	buf.Write([]byte("hellohello"))
	fmt.Println(buf.Len())
	var  retDataLen int64
	if err := binary.Read(&buf, binary.BigEndian, &retDataLen); err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf.Len())
}

