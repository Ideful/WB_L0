package main

import (
	qwe "encoding/json"
	"fmt"
	"l0/internal/generator"
	"l0/internal/models"
	repository "l0/internal/repository"
	"log"
	"time"
)

// func handleSTANMessage(data []byte) {
// 	// Реализуйте обработку сообщения, например, десериализацию и обработку данных
// 	fmt.Println("Received message:", string(data))
// }

// func subscribeToChannel(clusterID, clientID, channelName string) {
// 	// Подключение к NATS Streaming
// 	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://localhost:4222"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer sc.Close()

// 	// Подписка на канал
// 	sub, err := sc.Subscribe(channelName, func(msg *stan.Msg) {
// 		handleSTANMessage(msg.Data)
// 	}, stan.DurableName("durable-name"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer sub.Unsubscribe()

// 	// Ожидание завершения
// 	select {}
// }

// func publishToChannel(clusterID, clientID, channelName string, message []byte) {
// 	// Подключение к NATS Streaming
// 	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://localhost:4222"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer sc.Close()

// 	// Отправка сообщения в канал
// 	err = sc.Publish(channelName, message)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

const (
	nss       = "nats-streaming-server"
	nss_op    = "-sc"
	cfg_file  = "config/nats-streaming-config.yaml"
	port_flag = "-p"
)

func main() {
	// viper.SetConfigFile("config/nats-streaming-config.yaml")
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalf("Error reading config file: %v", err)
	// }

	// clusterID := viper.GetString("cluster_id")
	// clientID := viper.GetString("client_id")
	// channelName := viper.GetString("channel_name")

	// sc, err := stan.Connect(
	// 	clusterID,
	// 	clientID,
	// 	stan.NatsURL("nats://localhost:4222"),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer sc.Close()

	// go func() {
	// 	for {
	// 		val, err := generator.Generator()
	// 		if err != nil {
	// 			log.Println(err)
	// 		}

	// 		time.Sleep(50 * time.Millisecond)
	// 		err = sc.Publish(channelName, val)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		fmt.Println("Message sent successfully!")
	// 	}
	// }()

	// sub, err := sc.Subscribe(channelName, func(m *stan.Msg) {
	// 	fmt.Printf("Received a message: %s\n", string(m.Data))
	// }, stan.DurableName("my-durable"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer sub.Close()

	// for {
	// }

	db_cfg := repository.InitConfig()
	db, err := repository.CreatePostgresDB(db_cfg)

	if err != nil {
		fmt.Println(err)
	}
	var query = `INSERT INTO delivery_info (name, phone, zip, city, address, region,
				 email) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.DB.Exec(query, "a", "a", "a", "a", "a", "a", "a")
	if err != nil {
		fmt.Println(err)
	}

	for {
		json, err := generator.Generator()
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(string(json))
		var order models.Order
		err = qwe.Unmarshal(json, &order)
		delivery := order.Delivery
		// payment := order.Payment
		// items := order.Items

		fmt.Print(delivery)
		var query = `INSERT INTO delivery_info (name, phone, zip, city, address, region,
			email) VALUES ($1, $2, $3, $4, $5, $6, $7)`
		_, err = db.DB.Exec(query, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)

		time.Sleep(50 * time.Millisecond)

	}
	// db.Close()
}
