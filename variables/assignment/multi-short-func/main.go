package main

import f "fmt"

func main() {
	var (
		a int
		b int
	)

	_, b = multi()
	f.Println(b)
}

func multi() {
	return 5,4
}
