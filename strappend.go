package main

import (
	"time"
	"bytes"
	"fmt"
)
//效率：copy > append > buf.WriteString > += > fmt.Sprintf
//所以请慎用fmt.Sprinf 和 +=
func main() {
	str:="chinese"
	city:="beijing"

	// 1. +=
	s:=time.Now()
	for i:=0;i<100000;i++ {
		str +=city
	}
	e:=time.Since(s)
	fmt.Println("[+=]time cost 1:", e)

	// 2. fmt.Sprintf
	str="chinese"
	city="beijing"
	s=time.Now()
	for i:=0;i<100000;i++ {
		str = fmt.Sprintf("%s%s",str,city)
	}
	e=time.Since(s)
	fmt.Println("[fmt.Sprintf]time cost 2:", e)

	//3.  buffer.WriteString
	str="chinese"
	city="beijing"
	s=time.Now()
	var buf= bytes.Buffer{}
	buf.WriteString(str)
	for i:=0;i<100000;i++ {
		buf.WriteString(city)
	}
	e=time.Since(s)
	fmt.Println("[buffer.WriteString]time cost 3:", e)

	//4. append
	str="chinese"
	city="beijing"
	s=time.Now()
	bstr:=[]byte(str)
	bcity:=[]byte(city)
	for i:=0;i<100000;i++ {
		bstr= append(bstr, bcity...)
	}
	e=time.Since(s)
	fmt.Println("[append]time cost 4:", e)

	// 5. copy
	str="chinese"
	city="beijing"
	s=time.Now()
	zstr :=[]byte(str)
	for i:=0;i<100000;i++ {
		copy(zstr, city)
	}
	e=time.Since(s)
	fmt.Println("[copy]time cost 5:", e)
}
