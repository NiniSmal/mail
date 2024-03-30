package main

import (
	"fmt"
	"gitlab.com/nina8884807/mail/internal/api"
	gen "gitlab.com/nina8884807/mail/proto"
	"google.golang.org/grpc"
	"log"
	"net"
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
