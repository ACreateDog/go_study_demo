package main

import (
	"html/template"
	"path"
	"runtime"
)

func main() {
	p :=  getCurrentPath()
	p =  p + "/index.html"

	t , err := template.ParseFiles(p)
	if err != nil {

	}


}
func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}