package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatalf("Failed to listen on port 4040: ", err)
	}

	grcpServer := grpc.NewServer()

	if err := grcpServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve GRPC server on port 4040: ", err)
	}

}
