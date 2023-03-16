package main

import (
	"fmt"
	"net"
)

var startPort int
var endPort int
var i int
var ports int

func portrange() {
	fmt.Print("Enter startPort: ")
	fmt.Scanln(&startPort)

	fmt.Print("Enter endPort: ")
	fmt.Scanln(&endPort)
}

func main() {
	portrange()
	for i = startPort; i < endPort; i++ {
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
