package main

import "os"

func main() {
	fhandle , _  := os.Open("")
	fd := fhandle.Fd()


}

