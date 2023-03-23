package main

import (
	"fmt"
	"net"

	"C:\school\infra\S2\APC\SNIF\log.yaml"
)

type log struct {
	file string `yaml:"file"`
}

func portrange() (int, int) {
	var startPort int
	var endPort int

	fmt.Print("Enter startPort: ")
	fmt.Scanln(&startPort)

	fmt.Print("Enter endPort: ")
	fmt.Scanln(&endPort)
	return startPort, endPort
}

func portscan() {
	var i int
	var ports int

	startPort, endPort := portrange()
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

func main() {
	portscan()
}
