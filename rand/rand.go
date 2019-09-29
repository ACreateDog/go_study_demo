package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	frameIDInt := r.Intn(65535)
	fmt.Println(frameIDInt)
}
