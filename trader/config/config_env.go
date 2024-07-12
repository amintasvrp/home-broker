package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type ConfigEnv struct {
	BootstrapServers string
	GroupID          string
	Offset           string
}

func NewConfigEnv() *ConfigEnv {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	return &ConfigEnv{
		BootstrapServers: os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		GroupID:          os.Getenv("KAFKA_GROUP_ID"),
		Offset:           os.Getenv("KAFKA_OFFSET"),
	}
}
