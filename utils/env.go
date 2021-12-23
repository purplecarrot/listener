package utils

import (
	"log"
	"strconv"
	"strings"
)

func ProcessPortsEnv(port_strings []string) []int {
	var listen_ports []int

	for _, port_string := range port_strings {

		for _, chunk := range strings.Split(port_string, ",") {
			if len(chunk) > 0 {
				port_int, err := strconv.Atoi(chunk)
				if err != nil {
					log.Printf("discarding '%s' as it is not a valid port number\n", chunk)
				} else {
					listen_ports = append(listen_ports, port_int)
				}
			}
		}
	}

	return listen_ports
}
