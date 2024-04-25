package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"gitlab.com/nina8884807/mail/config"
	"gitlab.com/nina8884807/mail/internal/api"
	gen "gitlab.com/nina8884807/mail/proto"
	"gitlab.com/nina8884807/mail/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.Validation()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", cfg)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{cfg.KafkaAddr},
		Topic:     cfg.KafkaTopicCreateUser,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	defer r.Close()
	sendService := service.NewSendService(cfg.MailLogin, cfg.MailPassword)
	kafkaHandler := api.NewKafkaHandler(r, sendService)

	go kafkaHandler.OnCreateUser()

	log.Println("start grpc-server at: 8090")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	h := api.NewGrpcHandler(sendService)
	gen.RegisterMailServer(grpcServer, h)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}

// 1.сгенерировать ссылку с  уник. кодом, отправить его на почту.
//2. получить подтверждение от пользователя с этим кодом.
//3.Пометить пользователя , как верифицированного
