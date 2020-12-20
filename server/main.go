package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/siddontang/go-log/log"
	"io"
	"math/big"
	"net"
	"runtime/debug"
	"sync"
	"time"
)

var arrConn []net.Conn
var mapConnByte = make(map[net.Conn][]byte)
var lock sync.Mutex

func main() {
	debug.FreeOSMemory()
	return

	nowTimestamp := time.Now().Unix() // 秒级时间戳
	updateTime := "2020-12-18 12:04:38"
	oldTime, _ := time.ParseInLocation("2006-01-02 15:04:05", updateTime, time.Local)
	oldTimestamp := oldTime.Unix()
	fmt.Println(nowTimestamp)
	fmt.Println(oldTimestamp)
	a := nowTimestamp - oldTimestamp
	if a > 60 {
		fmt.Println("相隔60s")
	}

	return
	go randomWrite()
	go countConns()
	//go randomCloseConn()

	//listener , err  := net.Listen("tcp" , "0.0.0.0:1234")
	addr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:1234")
	listener, err := net.ListenTCP("tcp", addr)
	//fmt.Println("等待..连接...")
	if err != nil {
		log.Fatal("listener is error", err.Error())
		return
	}
	for {
		conn, err := listener.AcceptTCP()
		conn.SetKeepAlive(true)
		conn.SetKeepAlivePeriod(10 * time.Second)
		fmt.Printf("conn value  %+v", conn)
		if err != nil {
			fmt.Println("connection is error ", err.Error())
			continue
		}
		lock.Lock()
		arrConn = append(arrConn, conn)
		lock.Unlock()

		go process(conn)
	}
}
func BytesToInt(b []byte, i interface{}) interface{} {
	bytesBuffer := bytes.NewBuffer(b)
	binary.Read(bytesBuffer, binary.LittleEndian, i)
	return i
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var bufByte [8192]byte
		n, err := conn.Read(bufByte[:])
		if err != nil {
			fmt.Println("read is error connect is close", err.Error())
			if err == io.EOF {
				return
			}

		}
		fmt.Println("read data  is [ ", n, " ] bytes [ ", string(bufByte[:n]), " ]")
		mapConnByte[conn] = bufByte[:n]

		n, err = conn.Write(bufByte[:n])
		if err != nil {
			fmt.Println("write is error ", err)
		}
		fmt.Println("write data ", n)

	}
}
func randomWrite() {
	for {
		time.Sleep(2 * time.Second)

		if len(arrConn) <= 0 {
			continue
		}
		max := big.NewInt(int64(len(arrConn)))
		i, err := rand.Int(rand.Reader, max)

		if err != nil {
			fmt.Printf("randomWrite is error %v", err)
			continue
		}
		con := arrConn[i.Int64()]

		n, err := con.Write(mapConnByte[con])
		fmt.Printf("romdom write %v idxConn : %v \n", n, i.Int64())
		if err != nil {
			fmt.Printf("randomWrite con.Write is error %v \n", err)
			continue
		}
	}
}
func countConns() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("tcp long connection count is %v \n", len(arrConn))
	}
}
func randomCloseConn() {
	for {
		time.Sleep(10 * time.Second)
		max := big.NewInt(int64(len(arrConn)))
		i, err := rand.Int(rand.Reader, max)
		if err != nil {
			fmt.Printf("randomWrite is error %v", err)
			continue
		}
		lock.Lock()
		arrConn[i.Int64()].Close()
		arrConn = append(arrConn[:i.Int64()], arrConn[i.Int64()+1:]...)
		lock.Unlock()
	}
}
