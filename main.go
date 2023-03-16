package main

import (
	"fmt"
	"net"
)

func main() {
	var i uint
	ports := 0
	for i = 1; i < 100; i++ {
		addr := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", addr)
		if err == nil {
			fmt.Println(i, "port is open")
			conn.Close()
			ports = ports + 1
		} else {
			fmt.Println(i, "port is closed")
		}
	}

	fmt.Println("====================================")
	fmt.Println("====================================")
	fmt.Println("Scan is done")
	fmt.Println(ports, "ports are open")
	fmt.Println("====================================")
	fmt.Println("====================================")
}

// test change
