package autofit

import (
	"log"
	"net"
	"sync"
	"time"
)

var (
	i int64 = 0
	x int   = 0
	//Second int   = 0
	lock   sync.RWMutex
	err    error
	t      = time.Now()
	Second = t.Second()
)

func Bit62Adder(i int64) string {
	baseTable := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	a := string(baseTable[(i/62)%62])
	b := string(baseTable[(i/(62*62))%62])
	c := string(baseTable[(i/(62*62*62))%62])
	d := string(baseTable[(i/(62*62*62*62))%62])

	result := ""
	switch {
	case i < 62:
		result = string(baseTable[i%62])
	case i > 61 && i < 3844:
		result = a + string(baseTable[i%62])
	case i > 3843 && i < 238382:
		result = b + a + string(baseTable[i%62])
	case i > 238381 && i < 14776336:
		result = c + b + a + string(baseTable[i%62])
	case i > 14776335 && i < 916132822:
		result = d + c + b + a + string(baseTable[i%62])
	}
	return result
}

func TcpId(addr string) {
	t := time.Now()
	Second := t.Second()
	yearMap := make(map[int64]string)
	dateMap := make(map[int64]string)
	baseTable := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	yearTable := make([]int64, 62)
	n := int64(2021)
	for i := 0; i < 62; i++ {
		n = int64(2021 + i)
		yearTable[i] = n
	}
	for k, v := range baseTable {
		dateMap[int64(k)] = string(v)
		yearMap[int64(yearTable[k])] = string(v)
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error listener returned: %s", err)
	}
	defer l.Close()
	var lock sync.RWMutex
	for {
		Second = 0
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

				t := time.Now()
				if t.Second() > 58 {
					if x == 0 {
						Second = 0
					}
					x++
				}
				if t.Second() < 1 {
					x = 0
				}
				if t.Second() > Second {
					Second = t.Second()
					i = 0
				}
				_, err = c.Write([]byte(yearMap[int64(t.Year())] + dateMap[int64(t.Month())] + dateMap[int64(t.Day())] + dateMap[int64(t.Hour())] + dateMap[int64(t.Minute())] + dateMap[int64(t.Second())] + Bit62Adder(i)))
				lock.Lock()
				if t.Second() > Second {
					Second = t.Second()
					i = 0
				}
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
func GetId() string {
	yearMap := make(map[int64]string)
	dateMap := make(map[int64]string)
	baseTable := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	yearTable := make([]int64, 62)
	n := int64(2021)
	for i := 0; i < 62; i++ {
		n = int64(2021 + i)
		yearTable[i] = n
	}
	for k, v := range baseTable {
		dateMap[int64(k)] = string(v)
		yearMap[int64(yearTable[k])] = string(v)
	}
	t = time.Now()
	if t.Second() > 58 {
		if x == 0 {
			Second = 0
		}
		x++
	}
	if t.Second() < 1 {
		x = 0
	}
	if t.Second() > Second {
		Second = t.Second()
		i = 0
	}
	aaa := yearMap[int64(t.Year())] + dateMap[int64(t.Month())] + dateMap[int64(t.Day())] + dateMap[int64(t.Hour())] + dateMap[int64(t.Minute())] + dateMap[int64(t.Second())] + Bit62Adder(i)
	lock.Lock()
	if t.Second() < 1 {
		x = 0
	}
	i++
	lock.Unlock()
	return aaa
}
