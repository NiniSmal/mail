package api

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	gen "gitlab.com/nina8884807/mail/proto"
	"gitlab.com/nina8884807/mail/service"
	"log"
)

type KafkaHandler struct {
	service *service.SendService
	conn    *kafka.Reader
}

func NewKafkaHandler(conn *kafka.Reader, service *service.SendService) *KafkaHandler {
	return &KafkaHandler{
		service: service,
		conn:    conn,
	}
}

func (k *KafkaHandler) OnCreateUser() {
	for {
		mail, err := k.conn.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
		}
		var msg gen.SendEmailRequest

		err = json.Unmarshal(mail.Value, &msg)
		if err != nil {
			log.Println("parse message:", err)
			continue
		}
		err = k.service.SendMessage(&msg)
		if err != nil {
			log.Println("send message:", err)
			continue
		}
	}
}
