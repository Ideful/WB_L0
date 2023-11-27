package models

import (
	"log"

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

func (st *Stan) Disconnect() {
	if st.Sc != nil {
		st.Sc.Close()
	}
}
