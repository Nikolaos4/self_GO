package main 

import (
	"fmt"
	"bufio"
	"os"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(conn)
	for {
		fmt.Println("\nВведите строку: ")
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		fmt.Fprintf(conn, "%s\n", str)
		//
		if !scanner.Scan() {
			fmt.Println("Buy!")
			return
		}
		fmt.Printf("Текст: %s", scanner.Text(), )
	}
}
