package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var i int64 = 0

func main() {
	// Part 1: create a listener
	TcpId()
}
func TcpId() {
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	Hour := time.Now().Format("15")
	Minute := time.Now().Format("04")
	Second := time.Now().Format("05")

	//t.Format("15:04:05")
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Error listener returned: %s", err)
	}
	defer l.Close()
	var lock sync.RWMutex
	for {
		// Part 2: accept new connection
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("Error to accept new connection: %s", err)
		}

		// Part 3: create a goroutine that reads and write back data
		go func() {
			log.Printf("TCP session open")
			defer c.Close()

			for {
				d := make([]byte, 1024)

				// Read from TCP buffer
				_, err := c.Read(d)
				if err != nil {
					log.Printf("Error reading TCP session: %s", err)
					break
				}
				log.Printf("reading data from client: %s\n", string(d))

				// write back data to TCP client
				if time.Now().Format("2006") > year {
					year = time.Now().Format("2006")
					i = 0
				}
				if time.Now().Format("02") > day {
					day = time.Now().Format("02")
					i = 0
				}
				if time.Now().Format("15") > Hour {
					Hour = time.Now().Format("15")
					i = 0
				}
				if time.Now().Format("04") > Minute {
					Minute = time.Now().Format("04")
					i = 0
				}
				if time.Now().Format("05") > Second {
					Second = time.Now().Format("05")
					i = 0
				}
				_, err = c.Write([]byte(year[3:] + month + day + Hour + Minute + Second + fmt.Sprint(i)))
				lock.Lock()
				i++
				lock.Unlock()
				if err != nil {
					log.Printf("Error writing TCP session: %s", err)
					break
				}
			}
		}()

		// Part 4: create a goroutine that closes TCP session after 10 seconds
		go func() {
			// SetLinger(0) to force close the connection
			err := c.(*net.TCPConn).SetLinger(0)
			if err != nil {
				log.Printf("Error when setting linger: %s", err)
			}

			<-time.After(time.Duration(10) * time.Second)
			defer c.Close()
		}()
	}
}
