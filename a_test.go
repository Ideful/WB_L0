package test

import (
	"encoding/json"
	"fmt"
	nats "l0/internal/NATS"
	cache "l0/internal/cache"
	"l0/internal/models"
	"l0/internal/repository"
	"testing"
	"time"

	"github.com/nats-io/stan.go"
)

func TestDBCache(t *testing.T) {
	var st nats.Stan // соединяемся с nats-streaming
	if err := st.Connect(); err != nil {
		t.Fatal(err)
	}
	defer st.Disconnect()

	db, err := repository.CreatePostgresDB() // создаем и соединяемся с БД
	if err != nil {
		t.Fatal(err)
	}

	cache := cache.NewCache() // создаем кэш
	err = cache.FillCache(db) // заполняем данными из БД
	if err != nil {
		t.Fatal(err)
	}
	subscriber := nats.NewSubscriber(&st)       // создаем подписчика
	sub, err := subscriber.Subscribe(db, cache) // делаем подписку, а также внутри обновляем БД и кэш
	if err != nil {
		t.Fatal(err)
	}
	defer sub.Close()
	defer db.Db.Close()
	time.Sleep(1 * time.Second)
}

func TestPublisher(t *testing.T) {
	var st nats.Stan // соединяемся с nats-streaming
	if err := st.Connect(); err != nil {
		t.Fatal(err)
	}
	publisher := nats.NewPublisher(&st)
	go publisher.Publish()

	sub, _ := st.Sc.Subscribe("my-unique-channel", func(m *stan.Msg) {
		order := models.Order{}
		if err := json.Unmarshal(m.Data, &order); err != nil {
			t.Fatal(err)
		}
		if ok := models.Valid(order); !ok {
			fmt.Println("order file invalid")
		}
	}, stan.DurableName("Test"))
	defer sub.Close()
	time.Sleep(11 * time.Second)
}
