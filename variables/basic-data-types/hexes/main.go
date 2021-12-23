package main

import f "fmt"

func main(){
	for i := 0; i<16;i++ {
		f.Printf("waarde: 0x%x\n", i)
	}
	f.Printf("waarde: 0x%x\n", 17)
	f.Printf("waarde: 0x%x\n", 25)
	f.Printf("waarde: 0x%x\n", 50)
	f.Printf("waarde: 0x%x\n", 100)
}
