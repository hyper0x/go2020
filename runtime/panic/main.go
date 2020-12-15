package main

import "fmt"

type errDetail struct {
	msg    string
	number uint64
	point  float32
}

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Recovered Panic: %v\n", p)
		}
	}()
	panic(errDetail{
		msg:    "something wrong",
		number: 20210101,
		point:  92.7,
	})
}
