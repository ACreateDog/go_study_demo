package main

import (

	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const StateFllowerNode  = 0
const StateCandidateNode  = 1
const StateMasterNode  = 2

var state = StateFllowerNode
var AddressList = []string{
	"127.0.0.1:8087",
}
var r * rand.Rand
func init() {
	r = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
}
type DB struct {
	dBMap sync.Map
}
var db DB
var currentAddress = "127.0.0.1:8087"
var heartbeatChannel = make(chan string , 1)
var currentVoteIndex = 0
func main() {
	router := httprouter.New()

	router.GET("/set/:key/:value" , handleSet)
	router.GET("/get/:key" , handleGet)
	router.GET("/heartbeat" , handleHeartbeat)
	router.GET("/vote/:voteTarget/:voteIndex" , vote )
	router.GET("/ping" , ping)
	go  heandleResetHeartbeatTime()
	log.Fatal(http.ListenAndServe(currentAddress, router))
}
func vote(w http.ResponseWriter  , r * http.Request , ps httprouter.Params)  {
	voteTarget := ps.ByName("voteTarget")
	voteIndex := ps.ByName("voteIndex")
	intVoteIndex , _ := strconv.Atoi(voteIndex)
	if intVoteIndex > currentVoteIndex {
		db.Set("voteTarget",voteTarget)
		fmt.Fprint(w,"ok")
		currentVoteIndex = intVoteIndex
	} else {
		fmt.Fprint(w,"fail")
	}
}
func handleSet(w  http.ResponseWriter , r * http.Request , ps httprouter.Params) {

	key := ps.ByName("key")
	value := ps.ByName("value")

	db.Set(key, value)

	n , err :=  fmt.Fprintf(w , "ok,%s=%s" , key , value)
	if n <= 0 || err != nil {
		log.Println(err,"error exist here")
		return
	}

	return
}

func handleHeartbeat(w http.ResponseWriter , r *  http.Request , _ httprouter.Params) {
	n , err :=  io.WriteString(w,"ok")
	if n <= 0 || err != nil {
		log.Println(err)
	}
}
func ping(w http.ResponseWriter , r *  http.Request , _ httprouter.Params) {
	n , err :=  io.WriteString(w,"ok")
	if n <= 0 || err != nil {
		log.Println(err)
	}
}
func handleGet(w http.ResponseWriter , r * http.Request , ps httprouter.Params)  {
	key := ps.ByName("key")
	ret := db.Get(key)

	n , err := fmt.Fprint(w, "ok,",ret)
	if  n  <= 0 || err != nil {
		log.Fatal(err)
	}

	return
}

func heandleResetHeartbeatTime() {
	var t =  time.NewTimer(300 * time.Millisecond)

	for  {
		select {
		case <-heartbeatChannel:
			randInt := getRandomInt()
			d := time.Duration(randInt) * time.Millisecond
			t.Reset(d)

		case <-t.C:
			state = StateCandidateNode

		}

	}
}

func getRandomInt() int64 {
	return  int64((r.Intn(150) + 150))
}

func (d * DB)Get(key string) string {
	ret , ok :=  d.dBMap.Load(key)
	if ok {
		return ret.(string)
	}
	return ""
}
func (d *DB) Set(key string , value string)   {
	d.dBMap.Store(key , value)
}