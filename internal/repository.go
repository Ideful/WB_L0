package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

const (
	ordertalbe    = "orders"
	deliverytable = "delivery_info"
	paymenttable  = "payment_info"
	itemstable    = "order_items"
)

type Config struct {
	// Host     string `toml:"Host"`
	Port     string
	Username string
	Password string
	DBName   string
	// SSLMode  string `toml:"SSLMode"`
}

func InitConfig() Config {
	viper.SetConfigFile("config/db.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	port := viper.GetString("cluster_id")
	username := viper.GetString("client_id")
	password := viper.GetString("channel_name")
	dbname := viper.GetString("DBName")
	return Config{port, username, password, dbname}
}

func CreatePostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("port=%s user=%s dbname=%s password=%s ", cfg.Port,
		cfg.Username, cfg.DBName, cfg.Password))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
