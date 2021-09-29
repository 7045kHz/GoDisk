package main

import (
	"io"
	"log"
	"os"
)

func main() {
	var path string
	if len(os.Args) < 1 {
		path = "/dev/sda"
	} else {
		path = os.Args[1]
	}
	log.Println("[+] Reading boot secotr of " + path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}
	byteSlice := make([]byte, 512)
	n, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal("Error reading 512 bytes from file. " + err.Error())
	}
	//log.Printf("Bytes read: %d\n\n", n)
	/*
		myString :=string(byteSlice[:])
		log.Printf("Bytes dec: %d\n\n", byteSlice)
		log.Printf("Bytes oct: %o\n\n", byteSlice)
		log.Printf("Bytes hex: %x\n\n", byteSlice)
		log.Printf("Bytes string: %s\n\n", byteSlice)
		log.Printf("myString string: %s\n\n", myString)
		s:=""
		for _, v := range(byteSlice[:]) {
			if v > 0 && v <= 127 {
				s=s+string(v)
			}
		}

		log.Printf("%s" , s)
	*/
	log.Printf("%s", byteSlice[2:n])
	file.Close()

}
