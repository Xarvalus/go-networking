package main

import (
	"github.com/Xarvalus/go-networking/server/rpc/core"
	"github.com/Xarvalus/go-networking/server/utils"
	"log"
	"net/rpc"
)

func main() {
	tcpClient := connectViaTCP()
	httpClient := connectViaHTTP()

	tcpCall, replyTcp := fetchFiles(tcpClient, core.Args{CumulativelyUpTo: 2})
	httpCall, replyHttp := fetchFiles(httpClient, core.Args{CumulativelyUpTo: 3})

	select {
	case <-tcpCall.Done:
		log.Println("[TCP RPC]:", *replyTcp)
	}

	select {
	case <-httpCall.Done:
		log.Println("[HTTP RPC]:", *replyHttp)
	}

	defer closeConnections([]*rpc.Client{tcpClient, httpClient})
}

func connectViaTCP() *rpc.Client {
	client, err :=  rpc.Dial("tcp", utils.RpcPortTCP)
	utils.LogFatalError(err)

	return client
}

func connectViaHTTP() *rpc.Client {
	client, err :=  rpc.DialHTTP("tcp", utils.RpcPortHTTP)
	utils.LogFatalError(err)

	return client
}

func fetchFiles(client *rpc.Client, args core.Args) (*rpc.Call, *string) {
	var reply string

	call := client.Go("Data.FetchFiles", args, &reply, nil)

	return call, &reply
}

func closeConnections(clients []*rpc.Client) {
	for _, client := range clients {
		err := client.Close()
		utils.LogFatalError(err)
	}
}
