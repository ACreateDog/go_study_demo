package main

import (
	"bytes"
	"crypto/aes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
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
type MyInt int
func (m MyInt) String() string {
	return fmt.Sprint(int(m))   //BUG:死循环
}

func main() {
	//fooCopy()
	//nullJson()
	//zoreJons()
	//jsonS()
	//jsonZip()
	//jsonRegexp(getJsonStr())
	//jsonRegexp()
	//var a MyInt
	//a = 9
	//fmt.Println(a)

	tmp := map[string]interface{}{
		"a" : "a",
		"b" : 9,
		"c" : 9.8989,
	}
	for k , v := range tmp {
		fmt.Println(fmt.Sprintf("%s=%v",k,v))
	}
}
//  ([\]+)"([A-Za-z0-9_-]{1+})([\]+)":
func jsonRegexp( ) {
	//r :=  regexp.MustCompile("([\\\\]{0,})\"([A-Za-z0-9_-]{1,})([\\\\]{0,})\":")
	r :=  regexp.MustCompile("\"([A-Za-z0-9_-]{1,})\":")
	ret :=  r.FindAll([]byte(getJsonStr()) , -1)
	hashMap := make(map[string]string)
	var  sortList []string
	fmt.Println(len(ret))
	for i := 0 ; i < len(ret) ; i++  {

		dat := string(ret[i])
		if _ , ok := hashMap[dat]; !ok{
			hashMap[dat] = ""
			sortList  = append(sortList, dat)
		}
	}
	sort.Strings(sortList)
	fmt.Println(sortList)
	fmt.Println(len(sortList))
}
func jsonZip() {
	 m := make(map[string]interface{})
	 if err :=  json.Unmarshal([]byte(getJsonStr()) , &m) ; err != nil {
		fmt.Println(err)
	 }
	 // regex :=

	 //ret := make(map[string]string)
	for k , v := range m {
		switch v.(type) {
		case map[string]interface{}:
			fmt.Println("this a map = ", k, v )
			//tmpM := v.(map[string]interface{})
			//for kk , vv := range tmpM {
			//
			//}
		case string:
			fmt.Println("this a string : ", v )
		case float64:
			fmt.Println("this is a number : ", v )
		default:
			fmt.Println("not found type ")
		}
	}
}
func getJsonStrFromFile() []byte {
	filePath := "/Users/wangfei/Desktop/河图地图/河图1.4/全方案-1-hetu1.4.hetu"
	fd, err :=  os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	ret , err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println(err)
	}
	return ret
}
func getJsonStr() string  {
	return `
{
    "returnCode": 0,
    "returnMsg": "succ",
    "returnUserMsg": "成功",
    "data": {
        "0": [
            {
                "id": "1(0)",
                "attr": {
                    "deviceSN": "1",
                    "deviceType": 0,
                    "warehouseID": "297790069899526171",
                    "createdFrom": 1,
                    "mapName": "",
                    "mapVer": "",
                    "mapID": 1,
                    "curX": 3500,
                    "curY": 17500,
                    "theta": 1.570796,
                    "timestamp": 0,
                    "liftState": 1,
                    "liftHeight": 0,
                    "liftTheta": 0,
                    "frameId": "1",
                    "version": 0,
                    "packageID": "",
                    "generalIoInput": 0,
                    "generalIoOutput": 0,
                    "curDir": 0,
                    "payloadKg": 0,
                    "sideADir": 1,
                    "temper": 0,
                    "ucAlarm": 0,
                    "ucState": 0,
                    "ucStateSave": 0,
                    "ucStatus": 0,
                    "ucWork": 0,
                    "ulRunTime": 0,
                    "usSpeed": 0,
                    "voltage": 0,
                    "rotateModel": 0,
                    "seq": 1587029934167,
                    "ucPower": 100,
                    "state": 0,
                    "doorIsOpen": 0,
                    "status": 0,
                    "floor": "0",
                    "errorState": [],
                    "errorReason": 0,
                    "errorMessage": null,
                    "finishedOrderNumber": 0,
                    "finishedCodes": null,
                    "zoneID": "",
                    "pickLocationIndex": 0,
                    "placeLocationIndex": 0,
                    "mode": 0,
                    "curentSegment": 0,
                    "endX": 0,
                    "endY": 0,
                    "reserveStatus": 0
                },
                "workState": 3,
                "deviceID": "0",
                "deviceType": 0,
                "exceptionRules": []
            },
            {
                "id": "2(0)",
                "attr": {
                    "deviceSN": "2",
                    "deviceType": 0,
                    "warehouseID": "297790069899526171",
                    "createdFrom": 1,
                    "mapName": "",
                    "mapVer": "",
                    "mapID": 1,
                    "curX": 4500,
                    "curY": 14500,
                    "theta": 4.712389,
                    "timestamp": 0,
                    "liftState": 0,
                    "liftHeight": 0,
                    "liftTheta": 0,
                    "frameId": "65535",
                    "version": 0,
                    "packageID": "",
                    "generalIoInput": 0,
                    "generalIoOutput": 0,
                    "curDir": 0,
                    "payloadKg": 0,
                    "sideADir": 1,
                    "temper": 0,
                    "ucAlarm": 0,
                    "ucState": 0,
                    "ucStateSave": 0,
                    "ucStatus": 0,
                    "ucWork": 0,
                    "ulRunTime": 0,
                    "usSpeed": 0,
                    "voltage": 0,
                    "rotateModel": 0,
                    "seq": 1587030261819,
                    "ucPower": 100,
                    "state": 0,
                    "doorIsOpen": 0,
                    "status": 0,
                    "floor": "0",
                    "errorState": [],
                    "errorReason": 0,
                    "errorMessage": null,
                    "finishedOrderNumber": 0,
                    "finishedCodes": null,
                    "zoneID": "",
                    "pickLocationIndex": 0,
                    "placeLocationIndex": 0,
                    "mode": 0,
                    "curentSegment": 0,
                    "endX": 0,
                    "endY": 0,
                    "reserveStatus": 0
                },
                "workState": 3,
                "deviceID": "0",
                "deviceType": 0,
                "exceptionRules": []
            },
            {
                "id": "3003(0)",
                "attr": {
                    "deviceSN": "3003",
                    "deviceType": 0,
                    "warehouseID": "297790069899526171",
                    "createdFrom": 1,
                    "mapName": "",
                    "mapVer": "",
                    "mapID": 1,
                    "curX": 13500,
                    "curY": 14500,
                    "theta": 4.712389,
                    "timestamp": 0,
                    "liftState": 0,
                    "liftHeight": 0,
                    "liftTheta": 0,
                    "frameId": "65535",
                    "version": 0,
                    "packageID": "",
                    "generalIoInput": 0,
                    "generalIoOutput": 0,
                    "curDir": 0,
                    "payloadKg": 0,
                    "sideADir": 1,
                    "temper": 0,
                    "ucAlarm": 0,
                    "ucState": 0,
                    "ucStateSave": 0,
                    "ucStatus": 0,
                    "ucWork": 0,
                    "ulRunTime": 0,
                    "usSpeed": 0,
                    "voltage": 0,
                    "rotateModel": 0,
                    "seq": 1587116044712,
                    "ucPower": 100,
                    "state": 0,
                    "doorIsOpen": 0,
                    "status": 0,
                    "floor": "0",
                    "errorState": [],
                    "errorReason": 0,
                    "errorMessage": null,
                    "finishedOrderNumber": 0,
                    "finishedCodes": null,
                    "zoneID": "",
                    "pickLocationIndex": 0,
                    "placeLocationIndex": 0,
                    "mode": 0,
                    "curentSegment": 0,
                    "endX": 0,
                    "endY": 0,
                    "reserveStatus": 0
                },
                "workState": 7,
                "deviceID": "0",
                "deviceType": 0,
                "exceptionRules": []
            },
            {
                "id": "1379(0)",
                "attr": {
                    "deviceSN": "1379",
                    "deviceType": 0,
                    "warehouseID": "297790069899526171",
                    "createdFrom": 1,
                    "mapName": "",
                    "mapVer": "",
                    "mapID": 1,
                    "curX": 14500,
                    "curY": 22500,
                    "theta": 1.5707964,
                    "timestamp": 0,
                    "liftState": 0,
                    "liftHeight": 0,
                    "liftTheta": 0,
                    "frameId": "65535",
                    "version": 0,
                    "packageID": "",
                    "generalIoInput": 0,
                    "generalIoOutput": 0,
                    "curDir": 0,
                    "payloadKg": 0,
                    "sideADir": 1,
                    "temper": 0,
                    "ucAlarm": 0,
                    "ucState": 0,
                    "ucStateSave": 0,
                    "ucStatus": 0,
                    "ucWork": 0,
                    "ulRunTime": 0,
                    "usSpeed": 0,
                    "voltage": 0,
                    "rotateModel": 0,
                    "seq": 1587352989555,
                    "ucPower": 100,
                    "state": 0,
                    "doorIsOpen": 0,
                    "status": 0,
                    "floor": "0",
                    "errorState": [],
                    "errorReason": 0,
                    "errorMessage": null,
                    "finishedOrderNumber": 0,
                    "finishedCodes": null,
                    "zoneID": "",
                    "pickLocationIndex": 0,
                    "placeLocationIndex": 0,
                    "mode": 0,
                    "curentSegment": 0,
                    "endX": 0,
                    "endY": 0,
                    "reserveStatus": 0
                },
                "workState": 7,
                "deviceID": "0",
                "deviceType": 0,
                "exceptionRules": []
            }
        ],
        "1": [
            {
                "id": "9999(1)",
                "attr": {
                    "deviceSN": "9999",
                    "deviceType": 1,
                    "warehouseID": "",
                    "createdFrom": 0,
                    "mapName": "",
                    "mapVer": "",
                    "mapID": 0,
                    "curX": 0,
                    "curY": 0,
                    "theta": 1.5707964,
                    "timestamp": 0,
                    "liftState": 0,
                    "liftHeight": 0,
                    "liftTheta": 0,
                    "frameId": "",
                    "version": 0,
                    "packageID": "",
                    "generalIoInput": 0,
                    "generalIoOutput": 0,
                    "curDir": 0,
                    "payloadKg": 0,
                    "sideADir": 0,
                    "temper": 0,
                    "ucAlarm": 0,
                    "ucState": 0,
                    "ucStateSave": 0,
                    "ucStatus": 0,
                    "ucWork": 0,
                    "ulRunTime": 0,
                    "usSpeed": 0,
                    "voltage": 0,
                    "rotateModel": 0,
                    "seq": 0,
                    "ucPower": 0,
                    "state": 0,
                    "doorIsOpen": 0,
                    "status": 0,
                    "floor": "",
                    "errorState": [],
                    "errorReason": 0,
                    "errorMessage": null,
                    "finishedOrderNumber": 0,
                    "finishedCodes": null,
                    "zoneID": "156830853254",
                    "pickLocationIndex": 0,
                    "placeLocationIndex": 0,
                    "mode": 0,
                    "curentSegment": 0,
                    "endX": 0,
                    "endY": 0,
                    "reserveStatus": 0
                },
                "workState": 0,
                "deviceID": "1",
                "deviceType": 1,
                "exceptionRules": []
            }
        ]
    }
}
`
//	return `{
//"hetuConfig":"123938",
//"hetuConfig1":132435,
//"hetuConfig2" : {
//		"hetuConfig":8989,
//		"hetuConfig1":8989
//	}
//}`
}

func zoreJons() {
	var m   = make(map[string]interface{})
	m["deviceType"] = 0
	bts , _ := json.Marshal(m)
	fmt.Println(string(bts))
}
func jsonS()  {
	var a =  []Person{}
	bys , err := json.Marshal(a)
	if err != nil {

	}

	fmt.Println(string(bys))
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

//如果是空map转化成json，将变成{}
func nullJson() {
	var a = make(map[string]interface{})
	jsonStr , err  := json.Marshal(a)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(jsonStr))
}
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func EcbDecrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS7UnPadding(decrypted)
}

func EcbEncrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	data = PKCS7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}