package nats

import (
	"encoding/json"
	"fmt"
	cache "l0/internal/cache"
	models "l0/internal/models"
	"l0/internal/repository"
	"log"
	"time"

	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
)

type Config struct {
	ClusterID   string
	ClientID    string
	ChannelName string
}

type Stan struct {
	Sc  stan.Conn
	Cfg Config
}

type Publisher struct {
	publisher *Stan
}

func NewPublisher(s *Stan) *Publisher {
	return &Publisher{
		publisher: s,
	}
}

type Subscriber struct {
	subscriber *Stan
}

func NewSubscriber(s *Stan) *Subscriber {
	return &Subscriber{
		subscriber: s,
	}
}

func (st *Stan) InitConfig() {
	viper.SetConfigFile("config/nats-streaming-config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	st.Cfg.ClusterID = viper.GetString("cluster_id")
	st.Cfg.ClientID = viper.GetString("client_id")
	st.Cfg.ChannelName = viper.GetString("channel_name")
}

func (st *Stan) Connect() error {
	st.InitConfig()
	var err error
	st.Sc, err = stan.Connect(
		st.Cfg.ClusterID,
		st.Cfg.ClientID,
		stan.NatsURL("nats://localhost:4223"),
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *Publisher) Publish() {
	for {
		val, err := models.Generator()
		if err != nil {
			log.Println(err)
		}
		if err = p.publisher.Sc.Publish(p.publisher.Cfg.ChannelName, val); err != nil {
			log.Println(err)
		}
		time.Sleep(20 * time.Millisecond)
	}
}

func (s *Subscriber) Subscribe(db *repository.MyDB, c *cache.Cache) (stan.Subscription, error) {
	sub, err := s.subscriber.Sc.Subscribe(s.subscriber.Cfg.ChannelName, func(m *stan.Msg) {
		order := models.Order{}
		if err := json.Unmarshal(m.Data, &order); err != nil {
			fmt.Printf("json nvalid %v", err)
		}
		if ok := models.Valid(order); !ok {
			fmt.Println("order file invalid")
		}
		c.AddToCache(&order)
		if err := db.InsertQuery(&order); err != nil {
			log.Println(err)
		}

	}, stan.DurableName("test"))
	return sub, err
}

func (st *Stan) Disconnect() {
	if st.Sc != nil {
		st.Sc.Close()
	}
}
