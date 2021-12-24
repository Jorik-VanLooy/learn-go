package main

import (
	f "fmt"
	o "os"
)

func main() {
	name:=o.Args[1]

	f.Println("hi", name)
	f.Println("How are you?")
}
