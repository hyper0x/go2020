package main

import (
	"fmt"
	"hash"
	"hash/maphash"
)

func main() {
	var mh maphash.Hash
	var ok bool
	mhp := &mh
	if _, ok = interface{}(mhp).(hash.Hash); ok {
		fmt.Printf("%T implements interface hash.Hash.\n", mhp)
	} else {
		fmt.Printf("%T does not implement interface hash.Hash.\n", mhp)
	}
	if _, ok = interface{}(mhp).(hash.Hash32); ok {
		fmt.Printf("%T implements interface hash.Hash32.\n", mhp)
	} else {
		fmt.Printf("%T does not implement interface hash.Hash32.\n", mhp)
	}
	if _, ok = interface{}(mhp).(hash.Hash64); ok {
		fmt.Printf("%T implements interface hash.Hash64.\n", mhp)
	} else {
		fmt.Printf("%T does not implement interface hash.Hash64. \n", mhp)
	}
}
