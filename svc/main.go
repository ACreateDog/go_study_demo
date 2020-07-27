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

func (program) Init (env svc.Environment) error  {
	fmt.Println("program init...")
	return nil
}

func (this *program) Start() error {
	fmt.Println("start program")
	fmt.Println(syscall.Getpid())
	go func() {
		t := time.NewTicker( 2 * time.Second)
		this.t = t
		for d := range t.C {
			fmt.Println("tick at " , d)
		}
	}()
	return nil
}

func (this *program) Stop() error {
	this.t.Stop()
	fmt.Println("program stop ....")
	return nil
}

func main() {
	p := &program{}
	if err :=  svc.Run(p , syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM , syscall.SIGTERM,syscall.SIGKILL,syscall.SIGTSTP) ; err != nil {
		log.Fatal(err)
	}
}
