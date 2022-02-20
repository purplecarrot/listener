package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	network "purplecarrot/listen/network"
	"time"
)

type Port struct {
	Port  int    `json:"port"`
	Proto string `json:"proto"`
	Count int    `json:"use_count"`
}

func main() {

	arg_tcp_ports := flag.String("t", "", "TCP ports to listen on (multiple ports allowed separated by ,)")
	arg_udp_ports := flag.String("u", "", "UDP ports to listen on (multiple ports allowed separated by ,)")
	arg_status := flag.Int("s", 60, "print status every x seconds")
	flag.Parse()

	listenerPorts := make([]Port, 0)

	for _, tcp_port := range network.ProcessFlagString(*arg_tcp_ports) {
		port := &Port{
			Proto: "tcp",
			Port:  int(tcp_port),
			Count: 0,
		}
		listenerPorts = append(listenerPorts, *port)
	}

	for _, udp_port := range network.ProcessFlagString(*arg_udp_ports) {
		port := &Port{
			Proto: "udp",
			Port:  int(udp_port),
			Count: 0,
		}
		listenerPorts = append(listenerPorts, *port)
	}

	log.Printf("IP address is %v", network.GetIP())

	http.HandleFunc("/", network.HTTPHandler)
	for _, port := range listenerPorts {
		if port.Proto == "tcp" {
			go network.TCPListen(port.Port)
		}
		if port.Proto == "udp" {
			go network.UDPListen(port.Port)
		}
	}

	for {
		if *arg_status > 0 {
			log.Printf("[pid %d] alive with listeners on %v", os.Getpid(), listenerPorts)
			time.Sleep(time.Duration(*arg_status * int(time.Second)))
		}
	}

}
