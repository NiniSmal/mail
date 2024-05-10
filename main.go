package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"gitlab.com/nina8884807/mail/config"
	"gitlab.com/nina8884807/mail/internal/api"
	gen "gitlab.com/nina8884807/mail/proto"
	"gitlab.com/nina8884807/mail/service"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.GetConfig()
	if err != nil {
		l.Error("get config:", "error", err)
		return
	}

	err = cfg.Validation()
	if err != nil {
		l.Error("validation:", "error", err)
		return
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{cfg.KafkaAddr},
		Topic:     cfg.KafkaTopicCreateUser,
		GroupID:   "tm",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	defer r.Close()
	sendService := service.NewSendService(cfg.MailLogin, cfg.MailPassword)
	kafkaHandler := api.NewKafkaHandler(r, sendService, l)

	go kafkaHandler.OnCreateUser()

	l.Info(fmt.Sprintf("start grpc-server at: %d", cfg.Port))

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Port))
	if err != nil {
		l.Error("failed to listen", "error", err)
		return
	}
	grpcServer := grpc.NewServer()

	h := api.NewGrpcHandler(sendService)
	gen.RegisterMailServer(grpcServer, h)
	err = grpcServer.Serve(lis)
	if err != nil {
		l.Error("error", err)
	}

}

// 1.сгенерировать ссылку с  уник. кодом, отправить его на почту.
//2. получить подтверждение от пользователя с этим кодом.
//3.Пометить пользователя , как верифицированного
