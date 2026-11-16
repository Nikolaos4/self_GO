package main
import (
	"fmt"
	"net"
)

func main() {
	message := "Hello, I am a server"
	l, err := net.Listen("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()
	fmt.Println("Server is listening")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(message))
		conn.Close()
	}
}