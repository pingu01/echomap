package utils

import (
	"strconv"
	"strings"
)

func ParsePorts(ports string) []int {
	portSlice := make([]int, 0)

	for _, p := range strings.Split(ports, ",") {
		if strings.Contains(p, "-") {
			rangeParts := strings.Split(p, "-")
			if len(rangeParts) == 2 {
				startPort, err1 := strconv.Atoi(strings.TrimSpace(rangeParts[0]))
				endPort, err2 := strconv.Atoi(strings.TrimSpace(rangeParts[1]))
				if err1 == nil && err2 == nil && startPort <= endPort {
					// Add all ports within the range to the slice
					for port := startPort; port <= endPort; port++ {
						portSlice = append(portSlice, port)
					}
				}
			}
		} else {
			if port, err := strconv.Atoi(strings.TrimSpace(p)); err == nil {
				portSlice = append(portSlice, port)
			}
		}
	}
	return portSlice
}
