# go
##A time-based id generator
##征集替代秒后那个累加i60位进制字符串累加函数，要求一个60个字符串map通过轮盘累加原理实现的函数，进一步缩短编码长度。
##go get github.com/autofit/go/autofit

##Server Demo



`<package main

import (
	"github.com/autofit/go/autofit"
)

func main() {
	autofit.TcpId("9090")
}>`

##Client Demo:



`<package main

import (
	"log"
	"net"
)

func main() {
	c, err := net.Dial("tcp", ":3390")
	if err != nil {
		log.Fatalf("Error to open TCP connection: %s", err)
	}
	defer c.Close()
	// Part2: write some data to server
	log.Printf("TCP session open\t", c)
	b := []byte("0")

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
>`
