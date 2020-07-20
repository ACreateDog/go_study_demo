package main

import (
	"fmt"
	"github.com/withlin/canal-go/client"
	"log"
	"os"
	"time"
)

func main() {
	connector := client.NewSimpleCanalConnector("127.0.0.1", 11111, "canal", "canal", "example", 60000, 60*60*1000)
	err :=connector.Connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = connector.Subscribe(".*\\..*")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for {

		message,err := connector.Get(100, nil, nil)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		batchId := message.Id

		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("===没有数据了===")
			continue
		}

		fmt.Println(message)

	}
}
