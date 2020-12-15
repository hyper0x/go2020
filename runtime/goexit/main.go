package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		sec := time.Duration(3)
		fmt.Printf("Sleep %d second(s)...\n\n", sec)
		time.Sleep(time.Second * sec)
	}()
	runtime.Goexit()
}
