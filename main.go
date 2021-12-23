package main

import (
	"flag"
	"log"
	"os"
	"purplecarrot/listen/utils"
)

func main() {

	arg_tcp_ports := flag.String("t", "", "TCP ports to listen on (multiple ports allowed separated by ,)")
	arg_udp_ports := flag.String("u", "", "UDP ports to listen on (multiple ports allowed separated by ,)")
	flag.Parse()

	var listen_ports []int

	myIP := utils.GetIP()
	log.Printf("IP=%v", myIP)

	s := make([]string, 3)
	s = append(s, "2001")
	s = append(s, *arg_tcp_ports, *arg_udp_ports)
	s = append(s, os.Getenv("TCP_PORTS"), os.Getenv("UDP_PORTS"))

	listen_ports = utils.ProcessPortsEnv(s)
	log.Println("listening on ports", listen_ports)

}
