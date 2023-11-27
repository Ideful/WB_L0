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
)

func Publish(sc stan.Conn, cfg models.Config) {
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
}

func main() {
	var st models.Stan
	st.InitConfig()
	if err := st.Connect(); err != nil {
		log.Fatal(err)
	}
	defer st.Disconnect()

	go Publish(st.Sc, st.Cfg)

	db, err := repository.CreatePostgresDB(repository.InitConfig())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Db.Close()

	sub, err := st.Sc.Subscribe(st.Cfg.ChannelName, func(m *stan.Msg) {
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
