package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
	"io/fs"
)

type Config struct {
	Port                 int    `env:"PORT"`
	KafkaAddr            string `env:"KAFKA_ADDR"`
	KafkaTopicCreateUser string `env:"KAFKA_TOPIC_CREATE_USER"`
	MailLogin            string `env:"MAIL_LOGIN"`
	MailPassword         string `env:"MAIL_PASSWORD"`
}

func GetConfig() (*Config, error) {
	config := Config{}
	err := cleanenv.ReadConfig(".env", &config)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			err = cleanenv.ReadEnv(&config)
			if err != nil {
				return nil, err
			}
			return &config, nil
		}
		return nil, err
	}
	return &config, nil
}

func (c *Config) Validation() error {
	if c.Port == 0 {
		c.Port = 8090 //этот порт должен совпадать с открытым портом  контейнера
	}
	if c.KafkaAddr == "" {
		return errors.New("kafka address is empty")
	}
	if c.KafkaTopicCreateUser == "" {
		return errors.New("kafka topic create user is empty")
	}
	if c.MailLogin == "" {
		return errors.New("mail login is empty")
	}
	if c.MailPassword == "" {
		return errors.New("mail password is empty")
	}
	return nil
}
