package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/pingu01/echomap/cmd/scanner"
	"github.com/pingu01/echomap/utils"
)

var defaultPorts = []int{20, 21, 22, 23, 25, 53, 80, 110, 443, 445, 1433, 3306, 3389, 8080}
 
func main() {
	var host string
	var ports string

	flag.StringVar(&host, "t", "", "Hostname or IP to scan")
	flag.StringVar(&ports, "p", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(defaultPorts)),","), "[]"), "Ports to scan")

	flag.Usage = func() {
		fmt.Println("Usage: echomap -t=<hostname> -p=<port1,port2,... or range>")

		fmt.Println("OPTIONS:")
		fmt.Println("  -t :Target spefications, can be a hostname or an IP address")
		fmt.Println("  -p :Port specifications, can be a single port, multiple ports separated by commas, or a range of ports (e.g., 80,443,8080-8090)\n")

		fmt.Print("")

    }

	flag.Parse()
	
	if ports == "" {
		ports = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(defaultPorts)),","), "[]")
	}

	// if there are no host, print the usage
    if host == "" {
        flag.Usage()
        return
    }

	portSlice := utils.ParsePorts(ports)

	scanner.ScanPorts(host, portSlice)
}