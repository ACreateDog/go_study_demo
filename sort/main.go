package main

import (
	"fmt"
	"sort"
	"time"
)

type IntList []int

func (l*IntList) Len() int {
	return len(*l)
}

func (l*IntList) Less(i, j int) bool {
	return  (*l)[i] > (*l)[j]
}

func (l*IntList) Swap(i, j int) {
	(*l)[i] , (*l)[j] = (*l)[j] , (*l)[i]
}

func main() {
	var tmp IntList
	for i := 0; i < 100000 ; i ++  {
		tmp = append(tmp, i)
	}
	t := time.Now().UnixNano()
	sort.Sort(&tmp)
	t1 := time.Now().UnixNano()

	fmt.Println(t1 - t)

}
