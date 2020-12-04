package main

import (
	"fmt"
	"github.com/judwhite/go-svc/svc"
	"log"
	"syscall"
	"time"
)

type program struct {
	t *time.Ticker
}

func (program) Init(env svc.Environment) error {
	fmt.Println("program init...")
	return nil
}

func (this *program) Start() error {
	fmt.Println("start program")
	fmt.Println(syscall.Getpid())
	go func() {
		t := time.NewTicker(2 * time.Second)
		this.t = t
		for d := range t.C {
			fmt.Println("tick at ", d)
		}
	}()
	return nil
}

func (this *program) Stop() error {
	this.t.Stop()
	fmt.Println("program stop ....")
	return nil
}
func lengthOfLongestSubstringV1(s string) int {
	if len(s) <= 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	var chMap = make(map[string]int)
	for i := range s {
		k := string(s[i])
		if _, ok := chMap[k]; !ok {
			chMap[k] = 0
		}
		chMap[k]++
	}
	maxSubLength := len(chMap)
	minSubLength := 1
	max := 1
	step := 1
	for i := 0; i < len(s); i += step {
		var tmp = make(map[string]int)
		for window := minSubLength; window <= maxSubLength; window++ {
			idx := i + window - 1
			if idx >= len(s) {
				i = len(s)
				break
			}
			k := string(s[idx])
			if _, ok := tmp[k]; ok {
				step = tmp[k] + 1 - i
				break
			}
			tmp[k] = idx
		}
		if max < len(tmp) {
			max = len(tmp)
		}
	}

	return max
}
func lengthOfLongestSubstring(s string) int { //双指针滑动窗口
	max := 0
	n := len(s)
	i := 0
	hashmap := make(map[string]int)
	for j := 0; j < n; j++ {
		hashmap[string(s[j])]++
		for hashmap[string(s[j])] > 1 {
			i++
			hashmap[string(s[i])]--
		}
		if j-i+1 > max {
			max = j - i + 1
		}
	}
	return max
}
func main() {
	//var str = "aab"
	//fmt.Println(lengthOfLongestSubstring(str))
	t := time.Now().UnixNano()
	time.Sleep(1 * time.Second)
	t2 :=  time.Now().UnixNano()
	fmt.Println((t2 - t)/1e6)
	return
	p := &program{}
	if err := svc.Run(p, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGTSTP); err != nil {
		log.Fatal(err)
	}
}
