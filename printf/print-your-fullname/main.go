package main

import (
	f "fmt"
	o "os"
)
func main() {
	name, lastname := o.Args[1], o.Args[2]
	f.Printf("name: %s and lastname %s",name, lastname)
}
