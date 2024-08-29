package scanner

import (
    "fmt"
    "net"
    "sync"
    "time"
)

func scanPort(protocol, hostname string, port int) {
    address := fmt.Sprintf("%s:%d", hostname, port)
    conn, err := net.DialTimeout(protocol, address, 10*time.Second)
    if err != nil {
        return
    }
    conn.Close()
    fmt.Printf("%d - OPEN\n", port)
}

func ScanPorts(hostname string, ports []int) {
	fmt.Printf("Starting Echomap 0.1.0a at %s\n", time.Now().Format("2006-01-02 15:04:05"))
    var wg sync.WaitGroup
    for _, port := range ports {
        wg.Add(1)
        go func(p int) {
            defer wg.Done()
            scanPort("tcp", hostname, p)
        }(port)
    }
    wg.Wait()
}