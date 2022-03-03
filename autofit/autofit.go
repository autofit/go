package autofit

import (
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/robfig/cron"
)

var (
	Add int64 = 0
	X   int   = 0
	lock      sync.RWMutex
	err       error
	T         = time.Now()
	Second    = T.Second()
	yearMap   = make(map[int64]string)
	dateMap   = make(map[int64]string)
	baseTable = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	yearTable = make([]int64, 62)
	N         = int64(2021)
	RAND      = make([]string, 62)
)

func Bit62Adder(i int64) string {
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
	i := 0
	c := cron.New()
	spec := "*/1 * * * *"
	c.AddFunc(spec, func() {
		i++
		lock.Lock()
		Add = 0
		lock.Unlock()
	})
	c.Start()
	for i := 0; i < 62; i++ {
		N = int64(2021 + i)
		yearTable[i] = N
	}
	for k, v := range baseTable {
		dateMap[int64(k)] = string(v)
		yearMap[int64(yearTable[k])] = string(v)
		RAND[k] = string(v)
	}
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error listener returned: %s", err)
	}
	defer l.Close()
	for {
		Second = 0
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("Error to accept new connection: %s", err)
		}
		go func() {
			defer c.Close()
			for {
				T = time.Now()
				d := make([]byte, 1024)
				_, err := c.Read(d)
				if err != nil {
					log.Printf("Error reading TCP session: %s", err)
					break
				}
				r := make([]int, 4)
				for i := 0; i < 4; i++ {
					rand.Seed((time.Now().UnixNano()))
					r[i] = rand.Intn(62)

				}
				_, err = c.Write([]byte(yearMap[int64(T.Year())] + dateMap[int64(T.Month())] + dateMap[int64(T.Day())] + dateMap[int64(T.Hour())] + dateMap[int64(T.Minute())] + dateMap[int64(T.Second())] + Bit62Adder(Add) + ":" + dateMap[int64(T.Hour())+24] + dateMap[int64(T.Minute())] + dateMap[int64(T.Second())] + Bit62Adder(Add) + RAND[r[0]] + RAND[r[1]] + RAND[r[2]] + RAND[r[3]]))
				lock.Lock()
				Add++
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
func initAdd() {
	for {
		time.Sleep(1 * time.Second)
		T = time.Now()
		if T.Second() == 0 {
			lock.Lock()
			Add = 0
			lock.Unlock()

		}
	}
}
func GetId() string {

	for i := 0; i < 62; i++ {
		N = int64(2021 + i)
		yearTable[i] = N
	}
	for k, v := range baseTable {
		dateMap[int64(k)] = string(v)
		yearMap[int64(yearTable[k])] = string(v)
	}
	T = time.Now()
	if T.Second() > 58 {
		if X == 0 {
			Second = 0
		}
		X++
	}
	if T.Second() < 1 {
		X = 0
	}
	if T.Second() > Second {
		Second = T.Second()
		Add = 0
	}
	aaa := yearMap[int64(T.Year())] + dateMap[int64(T.Month())] + dateMap[int64(T.Day())] + dateMap[int64(T.Hour())] + dateMap[int64(T.Minute())] + dateMap[int64(T.Second())] + Bit62Adder(Add)
	lock.Lock()
	if T.Second() > Second {
		Second = T.Second()
		Add = 0
	}
	Add++
	lock.Unlock()
	return aaa
}
