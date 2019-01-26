package main

import (
	"github.com/Xarvalus/go-networking/server/rpc/core"
	"github.com/Xarvalus/go-networking/server/utils"
	"net"
	"net/http"
	"net/rpc"
	"sync"
)

// Simple RPC client via TCP & HTTP w/ concurrency
// TODO: gzip, encryption?
//
// Based on: "https://ops.tips/gists/example-go-rpc-client-and-server/"
func main() {
	server := rpc.NewServer()

	data := core.InitData()

	registerErr := server.Register(data)
	utils.LogFatalError(registerErr)

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	// Serves both HTTP and TCP handles asynchronously
	go serveViaTCP(server, &waitGroup)
	go serveViaHTTP(server, &waitGroup)

	waitGroup.Wait()
}

func serveViaTCP(server *rpc.Server, waitGroup *sync.WaitGroup) {
	listener, err := net.Listen("tcp4", utils.RpcPortTCP)
	utils.LogFatalError(err)

	server.Accept(listener)

	waitGroup.Done()
	defer listener.Close()
}

func serveViaHTTP(server *rpc.Server, waitGroup *sync.WaitGroup) {
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, err := net.Listen("tcp", utils.RpcPortHTTP)
	utils.LogFatalError(err)

	http.Serve(listener, nil)

	waitGroup.Done()
	defer listener.Close()
}
