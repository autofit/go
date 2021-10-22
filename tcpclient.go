package main

import (
	"log"
	"net"
)

func main() {
	c, err := net.Dial("tcp", "123.56.60.205:3390")
	if err != nil {
		log.Fatalf("Error to open TCP connection: %s", err)
	}
	defer c.Close()
	log.Printf("TCP session open\t", c)
	b := []byte("0")
	for i := 0; i < 100; i++ {

		_, err = c.Write(b)
		if err != nil {
			log.Fatalf("Error writing TCP session: %s", err)
		}
		d := make([]byte, 100)
		_, err = c.Read(d)
		if err != nil {
			log.Fatalf("Error reading TCP session: %s", err)
		}
		log.Printf("reading data from server: %s\n", string(d))
	}
}
