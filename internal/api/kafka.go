package api

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	gen "gitlab.com/nina8884807/mail/proto"
	"gitlab.com/nina8884807/mail/service"
	"log/slog"
)

type KafkaHandler struct {
	service *service.SendService
	conn    *kafka.Reader
	l       *slog.Logger
}

func NewKafkaHandler(conn *kafka.Reader, service *service.SendService, l *slog.Logger) *KafkaHandler {
	return &KafkaHandler{
		service: service,
		conn:    conn,
		l:       l,
	}
}

func (k *KafkaHandler) OnCreateUser() {
	for {
		mail, err := k.conn.ReadMessage(context.Background())
		if err != nil {
			k.l.Error("read message", "error", err)
			continue
		}
		var msg gen.SendEmailRequest

		err = json.Unmarshal(mail.Value, &msg)
		if err != nil {
			k.l.Error("parse message", "error", err)
			continue
		}
		err = k.service.SendMessage(&msg)
		if err != nil {
			k.l.Error("send message", "error", err)
			continue
		}
	}
}
