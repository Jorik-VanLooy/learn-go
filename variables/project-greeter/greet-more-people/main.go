package main

import (
	f "fmt"
	o "os"
)

func main() {
	count := len (o.Args) - 1
	f.Println("There are", count, "people !")
	for i:=1;i<=count;i++{
		f.Println("Hello great", o.Args[i],"!")
	}
	f.Println("Nice to meet you all")
}
