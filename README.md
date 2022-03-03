## A time-based Ratchet Algorithm Sequence Code Generator
## go get github.com/autofit/go/autofit
## Microcode platform：ai.cofjs.com
## Encoding server address based on ratcheting technology：ai.cofjs.com:53291
##### Platform development is simple：1.reg 2.create project 3.create table 
##### What you see is what you get.
##### The language translation system connects across languages.
##### Provide secondary development interface.
##### Provide privatized deployment package.
##### For Internet and Internet of Things development.
#### Update 62bit adder
### Server Demo & Client.Demo
### add a set of IoT data structure
### deviceshadow is a Highly abstract IoT data structure
### Iotdevinterface is a simplified version data interface base deviceshadow 
### lite is a lite version for iotdatastruck base on iotdevinterface.

```package main

import (
	"github.com/autofit/go/autofit"
)

func main() {
	autofit.TcpId("53291")
}

//client demo

package main

import (
	"log"
	"net"
)

func main() {
	c, err := net.Dial("tcp", ":53291")
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
	n, err := c.Read(d)
	if err != nil {
		log.Fatalf("Error reading TCP session: %s", err)
	}
	log.Printf("reading data from server: %s\n", string(d[:n]))
}
