package main

import f "fmt"

func main() {
	var (
		lang string
		version int
	)
	lang,version = "go",2
	f.Println(lang, "version", version)
}
