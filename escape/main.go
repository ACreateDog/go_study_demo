package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student) //局部变量s逃逸到堆

	s.Name = name
	s.Age = age

	return s
}
func Slice() {
	s := make([]int, 10000, 10000)

	for index, _ := range s {
		s[index] = index
	}
}
func Defunc() (foo string) {
	foo = "aaaa"
	defer func() {
		foo = "bbb"
	}()
	return
}
func longestPalindrome(s string) string {
	var bts = []byte(s)
	var ret string
	for i := 0 ; i < len(bts) ; i++ {
		tmp := center(bts, i)
		if len(ret) < len(tmp) {
			ret = tmp
		}
	}

	return ret
}

func center(bts []byte , startIndex int) string {
	if startIndex == 0 {
		return string(bts[0])
	}

	var left = startIndex
	var right = startIndex
	if left - 1 >= 0 && bts[left-1] == bts[startIndex]  {
		left--
	} else if right + 1 < len(bts) && bts[right+1] == bts[startIndex] {
		right++
	}
	for left >=0 && right < len(bts)  {

		if left - 1 >= 0  {
			left--
		}
		if  right + 1  < len(bts)  {
			right++
		}
	}
	return string(bts[left:right+1])
}
func main() {
	//StudentRegister("Jim", 18)
	//Slice()
	//fmt.Println(Defunc())
	//fmt.Println(time.Now().Add(10 * time.Second).UnixNano())
	//var t = 3

	//tmp := []byte("0123456789")
	//
	//fmt.Println(string(tmp[1:4]))
	fmt.Println(longestPalindrome("aaa"))
	return
	var a = "bbbb"
	defer func() {
		fmt.Println(a)
	}()
	defer func() {
		fmt.Println(a+"-hahah")
	}()
	a = "afafa"

	return
	//tmp , _ :=  json.Marshal(a)
	//fmt.Println(tmp)
	return
	num := 6
	for index := 0; index < num; index++ {
		resp, _ := http.Get("https://www.baidu.com")
		_, _ = ioutil.ReadAll(resp.Body)
	}
	fmt.Printf("此时goroutine个数= %d\n", runtime.NumGoroutine())
}
