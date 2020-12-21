package main

import "github.com/kataras/iris/v12"

func main() {
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
