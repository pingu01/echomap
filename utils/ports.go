package utils

import (
	"strconv"
	"strings"
)

func parsePorts(ports string) []int {
	portSlice := make([]int, 0)

	for _, p := range strings.Split(ports, ",") {
		if strings.Contains(p, "-") {
			rangePorts := strings.Split(p, "-")
			if len(rangePorts) == 2 {
				startPort, err1 := strconv.Atoi(rangePorts[0])
				endPort, err2 := strconv.Atoi(rangePorts[1])

				if err1 == nil && err2 == nil {
					for port := startPort; port <= endPort; port++ {
						portSlice = append(portSlice, port)
					}
				}
			}
		} else {
		// parse single port
			if port, err := strconv.Atoi(p); err == nil {
				portSlice = append(portSlice, port)
			}
		}
	}
	return portSlice
}