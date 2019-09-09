package main

import (
	"encoding/json"

	"fmt"
	"bytes"
	"time"
)

type Person struct {
	A int `json:"a"`
	Name string `json:"name"`
	Age int `json:"age"`
	Address struct{
		X int `json:"x"`
	} `json:"address"`
}
const (
	timeStandardLayout = "2006-01-02 15:04:05"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (u * User) MarshalJSON()(data []byte , err error) {
	type UserAlias User

	userStruct := &struct {
		CreateAt string `json:"create_at"`
		UpdateAt string `json:"update_at"`

	}{
		//UserAlias : (*UserAlias)(u),
	}
	userStruct.CreateAt = u.CreateAt.Format(timeStandardLayout)
	userStruct.UpdateAt = u.UpdateAt.Format(timeStandardLayout)

	return json.Marshal(userStruct)
}
func (u *User) UnmarshalJSON (data []byte)(err error) {
	type UserAlias  User

	return
}

func (p Person) UnmarshalJSON( b []byte) error {
	var buf bytes.Buffer
	 buf.Bytes()
	return nil
}
type Foo struct {
	Attr1 int `json:"attr1"`
	Attr1Struct interface{} `json:"attr1Struct"`
}
func main() {
	fooCopy()
}
func person() {
	var p Person
	bys := []byte("{\"a\":67,\"name\":\"tom\",\"age\":30 ,\"address\":\"beiJing 海淀\" }")
	bys = []byte("{\"a\":67,\"name\":\"tom\",\"age\":30}")

	err:= json.Unmarshal(bys , &p)
	if err != nil {
		fmt.Printf("error is %+v" , err)
		return
	}

	fmt.Printf("%+v" , p )
}
func fooSlice() {
	l1 := make([]byte , 50)
	l2 := make([]byte , 20)
	l1 = append(l1 , l2...)

	fmt.Println(len(l1))
}
func fooCopy() {
	l1 := make([]byte , 50)
	l2 := make([]byte , 20)
	copy(l1[40:] ,l2[:20] )
	handle(l1[25:] , l2[:10])
	fmt.Println(len(l1))

}
func handle(l1 , l2 []byte) {
	tmp := make([]byte , 10)
	l1 = append(l1 , tmp...)
	l2 = append(l2 , tmp...)

}
func foo() {
	var ff Foo
	jsonStr := "{\"attr1\":1,\"attr1Struct\":{\"f1\":2}}"
	jsonStr = "{\"attr1\":1}"
	err := json.Unmarshal([]byte(jsonStr) , &ff)
	if err != nil {
		fmt.Printf("Unmarshal is error %+v" , err)
		return
	}
	fmt.Println(ff.Attr1Struct == nil)
	//fmt.Println(IsNil(ff.Attr1Struct))
	fmt.Printf("Unmarshal is  %+v" , ff)
}
//func IsNil(i interface{}) bool {
//	vi := reflect.ValueOf(i)
//	fmt.Println(vi.Kind())
//	if vi.Kind() == reflect. {
//		return vi.IsNil()
//	}
//	return false
//}
