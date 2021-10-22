func TcpId() {
	t := time.Now()
	Second := t.Second()
	yearMap := make(map[int64]string)
	dateMap := make(map[int64]string)
	baseTable := []byte("123456789QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbn")
	yearTable := make([]int64, 60)
	n := int64(2021)
	for i := 0; i < 60; i++ {
		n = int64(2021 + i)
		yearTable[i] = n
	}
	for k, v := range baseTable {
		dateMap[int64(k)] = string(v)
		yearMap[int64(yearTable[k])] = string(v)
	}

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
				fmt.Println("i=", i, "现在的秒:", t.Second(), "过去的秒:", Second, "==", t.Second() > Second)
				if t.Second() > Second {
					if t.Second() == 0 {
						Second = 0
					}
					Second = t.Second()
					i = 0
				}
				_, err = c.Write([]byte(yearMap[int64(t.Year())] + dateMap[int64(t.Month())] + dateMap[int64(t.Day())] + dateMap[int64(t.Hour())] + dateMap[int64(t.Minute())] + dateMap[int64(t.Second())] + fmt.Sprint(i)))
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
