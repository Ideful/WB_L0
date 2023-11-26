package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	ordertalbe    = "orders"
	deliverytable = "delivery_info"
	paymenttable  = "payment_info"
	itemstable    = "order_items"
)

type Config struct {
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func InitConfig() Config {
	viper.SetConfigFile("config/db.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	port := viper.GetString("Port")
	username := viper.GetString("Username")
	password := viper.GetString("Password")
	dbname := viper.GetString("DBName")
	sslmode := viper.GetString("SSLMode")
	return Config{port, username, password, dbname, sslmode}
}

func CreatePostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Port,
		cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
