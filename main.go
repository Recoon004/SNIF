package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"gopkg.in/yaml.v2"
)

type Log struct {
	OpenPorts   []int `yaml:"open_ports"`
	ClosedPorts []int `yaml:"closed_ports"`
}

type Config struct {
	File string `yaml:"file"`
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

func portscan(cfg *Config) {
	var i int
	var openPorts []int
	var closedPorts []int

	startPort, endPort := portrange()
	for i = startPort; i < endPort; i++ {
		addr := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", addr)
		if err == nil {
			log.Println(i, "port is open")
			conn.Close()
			openPorts = append(openPorts, i)
		} else {
			log.Println(i, "port is closed")
			closedPorts = append(closedPorts, i)
		}
	}

	log.Println("====================================")
	log.Println("====================================")
	log.Println("Scan is done")
	log.Println(len(openPorts), "ports are open")
	log.Println("====================================")
	log.Println("====================================")

	logData := Log{
		OpenPorts:   openPorts,
		ClosedPorts: closedPorts,
	}

	logFile, err := os.Create(cfg.File)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}
	defer logFile.Close()

	encoder := yaml.NewEncoder(logFile)
	err = encoder.Encode(logData)
	if err != nil {
		log.Fatal("Failed to write log data: ", err)
	}
}

func main() {
	cfg := &Config{
		File: "log.yaml",
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	portscan(cfg)
}
