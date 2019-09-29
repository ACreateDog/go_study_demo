package common

import "sync"




//type FuncHandle func(data []byte)
//
//func ServerAt(addr string , handlerFunc FuncHandle) {
//	listener , err :=  net.Listen("tcp" , addr)
//	if err != nil {
//		log.Println("ServerAt,net.Listen,error=" , err)
//		return
//	}
//	for  {
//		conn , err :=  listener.Accept()
//		if err != nil {
//			log.Println("ServerAt,Accept,error=" , err)
//			continue
//		}
//
//		conn.Read()
//
//	}
//}
