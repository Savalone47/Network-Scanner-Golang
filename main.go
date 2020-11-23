package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)
func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}
	address := hostname + ":" + strconv.Itoa(port)
	conn,err := net.DialTimeout(protocol, address, 60 * time.Second)
	if err!= nil{
		result.State = "Closed"
		return result
	}
	defer conn.Close()
	result.State = "Open"
	return result
}

type ScanResult struct {
	Port string
	State string
}

func initialScan(hostname string) []ScanResult{
	var results []ScanResult
	for i:=1; i < 1024; i++{
		results = append(results, ScanPort("tcp", hostname,i))
		results = append(results, ScanPort("udp", hostname,i))
	}
	return results
}

func main() {
	fmt.Println("Port scanner in Go")

	open := ScanPort("tcp","localhost",1314)

	fmt.Printf("Port Open: %t\n",open)

	results := initialScan("localhost")
	fmt.Println(results)
}
