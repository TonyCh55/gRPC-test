package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Failed to listen on port 8080: ", err)
	}

	grcpServer := grpc.NewServer()

	if err := grcpServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve GRPC server on port 8080: ", err)
	}

}
