package nats

import (
	"encoding/json"
	"l0/internal/generator"
	"l0/internal/models"
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

func (st *Stan) Publish() {
	for {
		val, err := generator.Generator()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(50 * time.Millisecond)
		if err = st.Sc.Publish(st.Cfg.ChannelName, val); err != nil {
			log.Println(err)
		}
	}
}

func (st *Stan) Subscribe(db *repository.MyDB) (stan.Subscription, error) {
	sub, err := st.Sc.Subscribe(st.Cfg.ChannelName, func(m *stan.Msg) {
		var order models.Order
		if err := json.Unmarshal(m.Data, &order); err != nil {
			log.Println(err)

		}
		if err := db.ExecQuery(order); err != nil {
			log.Println(err)
		}

	}, stan.DurableName("nsame"))
	return sub, err
}

func (st *Stan) Disconnect() {
	if st.Sc != nil {
		st.Sc.Close()
	}
}