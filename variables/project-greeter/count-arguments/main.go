package main

import (
	f "fmt"
	o "os"
)

func main() {
	count:= len(o.Args) - 1

	f.Printf("there are %d arguments", count)
}
