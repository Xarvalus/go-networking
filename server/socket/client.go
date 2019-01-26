package main

import (
	"fmt"
	"github.com/Xarvalus/go-networking/server/utils"
	"log"
	"net"
)

func main() {
	tcpAddr, addrErr := net.ResolveTCPAddr("tcp4", utils.SocketPort)
	utils.LogFatalError(addrErr)

	conn, dialErr := net.DialTCP("tcp", nil, tcpAddr)
	utils.LogFatalError(dialErr)

	// Allows independent writes w/ reads
	go func() {
		for {
			input := utils.InputFromConsole()

			_, writeErr := conn.Write([]byte(input))
			if writeErr != nil {
				log.Println(writeErr)
				break
			}
		}
	}()

	request := make([]byte, 128)
	for {
		readLen, readErr := conn.Read(request)
		if readErr != nil {
			log.Println(readErr)
			break
		}

		if readLen > 0 {
			message := string(request[:readLen])
			fmt.Println("[WROTE]:", message)
		}
	}

	defer conn.Close()
}
