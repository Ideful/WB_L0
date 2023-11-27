package main

import (
	"encoding/json"
	"fmt"
	"l0/internal/generator"
	"l0/internal/models"
	"l0/internal/repository"
	"log"
	"time"

	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
)

type config struct {
	ClusterID   string
	ClientID    string
	ChannelName string
}

func (c *config) initConfig() {
	viper.SetConfigFile("config/nats-streaming-config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	c.ClusterID = viper.GetString("cluster_id")
	c.ClientID = viper.GetString("client_id")
	c.ChannelName = viper.GetString("channel_name")
}

func main() {
	var cfg config
	cfg.initConfig()
	sc, err := stan.Connect(
		cfg.ClusterID,
		cfg.ClientID,
		stan.NatsURL("nats://localhost:4223"),
	)
	if err != nil {
		log.Println(err)
	}
	defer sc.Close()

	go func() {
		for {
			val, err := generator.Generator()
			if err != nil {
				log.Println(err)
			}
			time.Sleep(50 * time.Millisecond)
			if err = sc.Publish(cfg.ChannelName, val); err != nil {
				log.Println(err)
			}
		}
	}()

	db, err := repository.CreatePostgresDB(repository.InitConfig())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Db.Close()

	sub, err := sc.Subscribe(cfg.ChannelName, func(m *stan.Msg) {
		var order models.Order
		if err := json.Unmarshal(m.Data, &order); err != nil {
			log.Println(err)
		}
		if err := db.ExecQuery(order); err != nil {
			log.Println(err)
		}
	}, stan.DurableName("nsame"))
	if err != nil {
		log.Println(err)
	}
	defer sub.Close()

	for {
	}
}
