package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"read-massage/internal/api"
	gen "read-massage/proto"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	gen.RegisterMailServer(grpcServer, &api.Handler{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

// 1.сгенерировать ссылку с  уник. кодом, отправить его на почту.
//2. получить подтверждение от пользователя с этим кодом.
//3.Пометить пользователя , как верифицированного
