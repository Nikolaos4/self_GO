package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8181")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Hello!")
	reader := bufio.NewReader(os.Stdin)
	scann := bufio.NewScanner(conn)
	for {
		fmt.Println("Введите строку: ")
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return 
		}
		if strings.TrimSpace(str) == "" {
			continue
		}
		fmt.Fprint(conn, str)
		//читаем
		if (!scann.Scan()) {
			fmt.Println("Buy")
			return
		}
		fmt.Println(scann.Text())
	}
}