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
	t := time.Now()
	year := t.Year()
	month := t.Month()
	day := t.Day()
	Hour := t.Hour()
	Minute := t.Minute()
	Second := t.Second()
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Error listener returned: %s", err)
	}
	defer l.Close()
	var lock sync.RWMutex
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("Error to accept new connection: %s", err)
		}
		go func() {
			log.Printf("TCP session open")
			defer c.Close()

			for {
				d := make([]byte, 1024)
				_, err := c.Read(d)
				if err != nil {
					log.Printf("Error reading TCP session: %s", err)
					break
				}
				log.Printf("reading data from client: %s\n", string(d))
				t := time.Now()
				if t.Year() > year {
					year = t.Year()
					i = 0
				}
				if t.Month() > month {
					month = t.Month()
					i = 0
				}
				if t.Day() > day {
					day = t.Day()
					i = 0
				}
				if t.Hour() > Hour {
					Hour = t.Hour()
					i = 0
				}
				if t.Minute() > Minute {
					Minute = t.Minute()
					i = 0
				}
				if t.Second() > Second {
					Second = t.Second()
					i = 0
				}
				_, err = c.Write([]byte(fmt.Sprint("%d", year)[3:] + fmt.Sprintf("%d", month) + fmt.Sprintf("%d", day) + fmt.Sprintf("%d", Hour) + fmt.Sprintf("%d", Minute) + fmt.Sprintf("%d", Second) + fmt.Sprint(i)))
				lock.Lock()
				i++
				lock.Unlock()
				if err != nil {
					log.Printf("Error writing TCP session: %s", err)
					break
				}
			}
		}()

		go func() {
			err := c.(*net.TCPConn).SetLinger(0)
			if err != nil {
				log.Printf("Error when setting linger: %s", err)
			}

			<-time.After(time.Duration(10) * time.Second)
			defer c.Close()
		}()
	}
}
