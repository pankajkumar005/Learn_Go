package main

import (
	"fmt"
	"example.com/user/hello/morestrings"
	"example.com/user/hello/prefixhello"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(prefixhello.PrefixHello("!oG ,olleH"))
}