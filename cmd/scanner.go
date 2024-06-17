package scanner

import(
	"fmt";
	"net";
	"sync";
	"time";
)

func scanPort(protocol, hostname string, port int){
	// this formats the address for better use later
	address := fmt.Sprintf("%s:%d", hostname, port)

	conn, err := net.DialTimeout(protocol, address, 10*time.Second)

	if err != nil{
		return
	}

	conn.Close()
	fmt.Printf("%d - OPEN\n", port)
}

func scanPorts(hostname string, ports []int){

	// wg now is a collection of goroutines, it is crucial for concurrent op's
	var wg sync.WaitGroup

	for _, port := range ports{
		wg.Add(1)
		go func(p int){
			// decreases the wg's counter when the goroutine completes
			defer wg.Done()
			//tcp scan only until now
			scanPort("tcp", hostname, p)
		}(port)
	}
	//ends the function when the wg is 0
	wg.Wait()
}
