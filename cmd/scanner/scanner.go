package scanner

import (
    "fmt"
    "net"
    "sync"
    "time"
)

func scanPort(protocol, hostname string, port int, verbose bool, results *[]int, mu *sync.Mutex) {
    address := fmt.Sprintf("%s:%d", hostname, port)
    conn, err := net.DialTimeout(protocol, address, 10*time.Second)
    if err != nil {
        return
    }
    conn.Close()

    if verbose {
        fmt.Printf("Discovered open port: %d/%s \n", port, protocol)
    }

    mu.Lock()
    *results = append(*results, port) 
    mu.Unlock()
}

func ScanPorts(hostname string, ports []int, verbose bool) {
    fmt.Printf("Starting Echomap 0.1.0a at %s\n", time.Now().Format("2006-01-02 15:04:05"))
    var wg sync.WaitGroup
    var mu sync.Mutex
    var results []int

    for _, port := range ports {
        wg.Add(1)
        go func(p int) {
            defer wg.Done()
            scanPort("tcp", hostname, p, verbose, &results, &mu)
        }(port)
    }
    wg.Wait()

    fmt.Println("\nScan completed. Open ports:")
    for _, port := range results {
        fmt.Printf("%d - OPEN\n", port)
    }
}
