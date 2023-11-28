package main

import (
	"fmt"

	nats "l0/internal/NATS"
	"l0/internal/repository"
	"log"
)

// func Subscribe

func main() {
	var st nats.Stan
	st.InitConfig()
	if err := st.Connect(); err != nil {
		log.Fatal(err)
	}
	defer st.Disconnect()

	db, err := repository.CreatePostgresDB(repository.InitConfig())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Db.Close()

	go st.Publish()

	sub, err := st.Subscribe(db)
	if err != nil {
		fmt.Println(err)
	}
	defer sub.Close()

	for {
	}
}
