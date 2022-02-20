package network

import (
	"log"
	"strconv"
	"strings"
)

func ProcessFlagString(flag string) []int {
	var ports []int

	for _, chunk := range strings.Split(flag, ",") {
		if len(chunk) > 0 {
			port_int, err := strconv.Atoi(chunk)
			if err != nil {
				log.Printf("discarding '%s' as it is not a valid port number\n", chunk)
			} else {
				ports = append(ports, port_int)
			}
		}
	}

	return ports
}
