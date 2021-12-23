package main

import f "fmt"

func main() {
	var (
		planet string
		isTrue bool
		temp float64
	)
	planet = "Mars"
	isTrue = true
	temp = 19.5

	f.Println ("Air is good on", planet,"\nIt's", isTrue, "\nIt is",temp,"degrees")
}
