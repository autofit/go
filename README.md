## A time-based Ratchet Algorithm Sequence Code Generator
## go get github.com/autofit/go/autofit
## Microcode platform：ai.cofjs.com:81
## Encoding server address based on ratcheting technology：ai.cofjs.com:9090
##### Platform development is simple：1.reg 2.create project 3.create table 
##### What you see is what you get.
##### The language translation system connects across languages.
##### Provide secondary development interface.
##### Provide privatized deployment package.
##### For Internet and Internet of Things development.
#### Update 62bit adder
### Server Demo & Client.Demo

```package main

import (
	"github.com/autofit/go/autofit"
)

func main() {
	autofit.TcpId("9090")
}

//client demo

package main

import (
	"log"
	"net"
)

func main() {
	c, err := net.Dial("tcp", ":9090")
	if err != nil {
		log.Fatalf("Error to open TCP connection: %s", err)
	}
	defer c.Close()
	log.Printf("TCP session open\t", c)
	_, err = c.Write([]byte("0")
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
