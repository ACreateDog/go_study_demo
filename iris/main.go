package main

import (
	"demo/iris/pkg"
	"fmt"
	"github.com/kataras/iris/v12"
	"time"
)

type Foo struct {
	Name string
}

func main() {
	var list []Foo

	list = append(list, Foo{"n1"})
	list = append(list, Foo{"n2"})
	list = append(list, Foo{"n3"})
	fmt.Println(list)
	var ret = make([]*Foo,3)
	for i,l := range list {
		//tmp := l
		ret[i] = &l
	}
	fmt.Println(ret)
	return

	t :=  time.Now()

	time.Sleep(2 * time.Second)

	foo := time.Now().Sub(t).Milliseconds()
	fmt.Println(foo)



	return
	fmt.Println(pkg.Lib)
	fmt.Println(pkg.Foo)
	//return
	app := iris.New()
	app.Get("/index",index)
	if err :=  app.Listen(":8888");err != nil {
		return
	}
}
func index(c iris.Context) {

	c.JSON(map[string]interface{}{
		"key":c.Params().Get("key"),
		"value":c.Params().Get("val"),
	})
	return
}
