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
	var verbose bool

	flag.StringVar(&host, "t", "", "Hostname or IP to scan")
	flag.StringVar(&ports, "p", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(defaultPorts)),","), "[]"), "Ports to scan")
	flag.BoolVar(&verbose, "v", false, "Verbose output")

	flag.Usage = func() {
		fmt.Println("Usage: echomap -t=<hostname> -p=<port1,port2,...>")

		fmt.Println("OPTIONS:")
		fmt.Println("  -t :Target spefications, can be a hostname or an IP address")
		fmt.Println("  -p :Port specifications, can be a single port or a range of ports separated by commas\n      Example: -p=80,443,8080")

		fmt.Print("")

    }

	// Parse the command-line flags
	flag.Parse()  

	if flag.NArg() > 0 {
		fmt.Println("Usage: echomap -t=<hostname> -p=<port1,port2,...>")
		return
	}

	if host == "" {
		fmt.Println(flag.Usage)
		return
	}

	portSlice := utils.ParsePorts(ports)

	scanner.ScanPorts(host, portSlice, verbose)
}