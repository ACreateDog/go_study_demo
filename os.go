package main

import (
	"os"
	"github.com/siddontang/go-log/log"
	"fmt"
)

func main() {

	foo , err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(foo)

}
