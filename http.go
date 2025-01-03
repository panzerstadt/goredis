package main

import (
	"fmt"
	"io"
	"net"
)

func serve_http(c chan int) {
	fmt.Println("starting up http protocol server at port 5000.")
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		c <- 1
	}

	fmt.Printf("new connection on %v\n", conn.RemoteAddr())

	defer conn.Close()

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	fmt.Println("sent back OK")

	for {
		packet := make([]byte, 4096)
		tmp := make([]byte, 4096)

		_, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			println("END OF FILE")
			break
		}
		packet = append(packet, tmp...)
		fmt.Println(string(packet))
	}
}
