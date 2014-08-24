package main

import (
	"fmt"
	"time"
)

func main() {
	a := make([]byte, 1024*1024*1024*3)
	for i := 0; i < len(a); i++ {
		a[i] = 0xff
	}
	fmt.Println("created")
	time.Sleep(5 * time.Second)
	for i := 1024 * 1024 * 1024; i < len(a); i++ {
		a[i] = 0
	}
	time.Sleep(1 * time.Minute)
	a[4] = 5
}
