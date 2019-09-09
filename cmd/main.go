package main

import (
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"fmt"
	"time"
	"sync"
	"net"
)

var mutex sync.Mutex

func main() {

	//mutex.Lock()
	//ch := make(chan bool)
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	log.Println("获取锁1...")
	//	mutex.Lock()
	//	log.Println("获取到锁1...")
	//}()
	//
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	log.Println("获取锁2...")
	//	mutex.Lock()
	//	log.Println("获取到锁2...")
	//}()
	//
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	log.Println("释放锁3...")
	//	mutex.Unlock()
	//}()

	//time.Sleep(30 * time.Second)
	//<-ch


	client , err :=  net.Dial("tcp" , "112.124.110.150:1234")
	//client , err :=  net.Dial("tcp" , "114.118.30.25:1234")
	if err != nil {
		log.Println("connect" ,err)
	}

	log.Printf("client %v" , client)

	client.Write([]byte("23&"))
	var data []byte
	n  , err := client.Read(data)
	log.Println(n)

}



func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func run() {
	go func() {
		time.Sleep(2 * time.Second)
		for i := 0; i < 10; i++ {
			Send(i)
		}
	}()
	Receive()
}
func Receive() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		//for d := range msgs {
		//	log.Printf("Received a message: %s", d.Body)
		//}
		for {
			select {
			case d := <-msgs:
				log.Printf("Received a message: %s", d.Body)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func Send(i int) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()


	ch , err := conn.Channel()
	failOnError(err , "failed to open a channel")
	defer ch.Close()


	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!" + strconv.Itoa(i)
	fmt.Println(body)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")


}