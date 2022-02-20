package network

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func GetIP() net.IP {
	conn, err := net.Dial("udp", "1.1.1.1:53")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func handleUDPConn(conn *net.UDPConn, connID string) {

	buffer := make([]byte, 512)

	n, remoteaddr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Printf("error reading from port %v - %v ", conn.LocalAddr().String(), err)
	}

	data := string(buffer[:n])

	log.Printf("[%s] >%s sent us %v bytes \"%s\"", connID, remoteaddr, len(data), strings.TrimSuffix(data, "\n"))

	response := fmt.Sprintf("%d received %d bytes from you at %s", os.Getpid(), len(data), remoteaddr)
	_, err = conn.WriteToUDP([]byte(response), remoteaddr)
	if err != nil {
		log.Printf("udp error - couldn't send UDP response %v", err)
	} else {
		log.Printf("[%s] <%s sent back \"%s\"", connID, remoteaddr, response)
	}

}

func UDPListen(port int) {

	log.Printf("[%d/udp] listening on UDP port %d", port, port)

	connID := fmt.Sprintf("%d/udp", port)

	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)

	if err != nil {
		log.Printf("[%d/udp] error binding to port - %v\n", port, err)
	}

	defer conn.Close()

	for {
		handleUDPConn(conn, connID)
	}

}
