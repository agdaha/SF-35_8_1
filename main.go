package main

import (
	"fmt"
	"log"
	"net"
	"skillfactory/SF-35_8_1/message"
	"time"
)

const addr = "0.0.0.0:3210"

const proto = "tcp4"

func main() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	log.Println("Start server")
	defer listener.Close()
	i := 1
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		log.Printf("Init connection#%v", i)
		go handleConnection(conn, i)
		i += 1
	}
}

func handleConnection(c net.Conn, i int) {
	defer func() {
		c.Close()
		log.Printf("Close connection#%v", i)
	}()
	for {
		_, err := c.Write([]byte(message.Get() + "\n"))
		if err != nil {
			log.Printf("Error write to connection#%v: %v", i, err)
			break
		}
		log.Printf("write to connection#%v", i)
		time.Sleep(time.Second * 3)
	}
}
