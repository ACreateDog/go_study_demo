package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"strconv"
	"time"
)

type Node struct {
	ID uint64 `gorm:"primary_key;column:id" json:"id"`
	NodeID string `gorm:"column:node_id" json:"nodeId"`
	WarehouseID string `gorm:"column:warehouse_id" json:"warehouseId"`
	X int `gorm:"column:x" json:"x"`
	Y int `gorm:"column:y" json:"y"`
	Z int `gorm:"column:z" json:"z"`
	MapID int `gorm:"column:map_id" json:"mapId"`
	LockStatus int `gorm:"column:lock_status" json:"lockStatus"`
	LockType int `gorm:"column:lock_type" json:"lockType"`
	Property string `gorm:"column:property" json:"property"`
	Status int `gorm:"column:status" json:"status"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`

}
func main() {
	r := rand.New( rand.NewSource(time.Now().Unix()))
	db , err := gorm.Open("mysql","root:@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i <= 150000000 ; i++ {
		node := Node{}
		node.NodeID = strconv.Itoa(r.Intn(50000000))
		node.WarehouseID = strconv.Itoa(r.Intn(50000000))
		node.MapID = r.Intn(50)
		node.X = r.Intn(500000)
		node.Y = r.Intn(500000)
		node.Z = 0
		node.CreateTime = time.Now()
		node.UpdateTime = time.Now()
		err =  db.Create(&node).Error
		if err != nil {
			fmt.Println(err)
		}

	}
	defer db.Close()
}



