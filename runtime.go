package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	os := runtime.GOOS
	today := time.Now()
	if os == "darwin" { os = "OS X" }
	fmt.Printf("go runs on %s\nand the time is now %s", os, today)
}