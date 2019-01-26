package main

import (
	"fmt"
	"github.com/Xarvalus/go-networking/server/utils"
	"log"
	"net"
)

// Simple chat Client <-> Server
//
// Based on: "https://astaxie.gitbooks.io/build-web-application-with-golang/en/08.1.html"
func main() {
	tcpAddr, addrErr := net.ResolveTCPAddr("tcp4", utils.SocketPort)
	utils.LogFatalError(addrErr)

	listener, listenErr := net.ListenTCP("tcp", tcpAddr)
	utils.LogFatalError(listenErr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)

			// do not exit the server
			continue
		}

		// By using 'goroutine' will handle multiple connections
		go manageSingleConnectionWrite(conn)
		go manageSingleConnectionRead(conn)
	}
}

func manageSingleConnectionWrite(conn net.Conn) {
	for {
		input := utils.InputFromConsole()

		_, writeErr := conn.Write([]byte(input))
		if writeErr != nil {
			log.Println(writeErr)
			break
		}
	}

	defer conn.Close()
}


func manageSingleConnectionRead(conn net.Conn) {
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
