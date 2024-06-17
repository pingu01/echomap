package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time";

	"github.com/pingu01/echomap/cmd/scanner"
)

var defaultPorts = []int{20, 21, 22, 23, 25, 53, 80, 110, 443, 445, 1433, 3306, 3389, 8080}
 
func main() {
	var host string
	var ports string

	flag.StringVar(&host, "t", "", "Hostname or IP to scan")
	flag.StringVar(&ports, "p", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(defaultPorts)),","), "[]"), "Ports to scan")

	flag.Usage = func() {
		fmt.Println("Usage: echomap -t=<hostname> -p=<port1,port2,...>")

		fmt.Println("OPTIONS:")
		fmt.Println("  -t :Target spefications, can be a hostname or an IP address")
		fmt.Println("  -p :Port specifications, can be a single port or a range of ports separated by commas\n      Example: -p=80,443,8080")

		fmt.Print("")

    }

	flag.Parse()  

	if flag.NArg() > 0 {
		fmt.Println("Usage: echomap -t=<hostname> -p=<port1,port2,...>")
		return
	}
	
	fmt.Printf("Starting Echomap 0.1.0a at %s\n", time.Now().Format("2006-01-02 15:04:05"))

    portSlice := make([]int, 0)
    for _, p := range strings.Split(ports, ",") {
        if port, err := strconv.Atoi(p); err == nil {
            portSlice = append(portSlice, port)
        }
    }

	scanner.scanPorts(host, portSlice)
}