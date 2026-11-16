package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		var word string
		_, err := fmt.Scan(&word)
		if err != nil {
			fmt.Println(err)
			continue
		}
		n, err := conn.Write([]byte(word)) 
		if err != nil || n == 0 {
			fmt.Println(err)
			return
		}
		fmt.Println("Перевод: ")
		conn.SetReadDeadline(time.Now().Add(time.Second*5))
		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(string(buff[:n]))
			conn.SetReadDeadline(time.Now().Add(time.Millisecond*700))
		}
		fmt.Println()
	}
}