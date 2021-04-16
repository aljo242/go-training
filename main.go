package main

import (
	"fmt"

	"github.com/aljo242/go-training/list"
)

func main() {
	fmt.Printf("Hello, world!\n")

	l := list.New()
	el := l.PushBack(1)
	el = l.InsertAfter(10, el)

	fmt.Println(el)
}
