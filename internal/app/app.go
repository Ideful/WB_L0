package app

import (
	"fmt"
	nats "l0/internal/NATS"
	cache "l0/internal/cache"
	models "l0/internal/models"
	"l0/internal/repository"
	"log"
	"net/http"
	"strconv"
)

var loginFormTmpl = []byte(`
<html>
	<body>
	<form action="/orders/" method="get">
		<input type="text" name="id">
		<input type="submit" value="id">
	</form>
	</body>
</html>
`)

func orders(w http.ResponseWriter, r *http.Request, db *repository.MyDB, cache *cache.Cache) {
	w.Write(loginFormTmpl)
	if r.Method != http.MethodGet {
		return
	}

	id := r.FormValue("id")
	if id == "" {
		return
	}
	val, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}

	o, err := cache.GetFromCache(val)
	if err != nil {
		w.Write([]byte("wrong idx"))
	}
	w.Write(o)
}

func Run() {
	var st nats.Stan // соединяемся с nats-streaming
	if err := st.Connect(); err != nil {
		log.Fatal(err)
	}
	defer st.Disconnect()

	publisher := nats.NewPublisher(&st) // создаем сущность публикующего
	go publisher.Publish()              // публикуем данные в канал

	db, err := repository.CreatePostgresDB() // создаем и соединяемся с БД
	if err != nil {
		fmt.Println(err)
	}
	defer db.Db.Close()

	cache := cache.NewCache() // создаем кэш
	err = cache.FillCache(db) // заполняем данными из БД
	if err != nil {
		fmt.Println(err)
	}

	subscriber := nats.NewSubscriber(&st)       // создаем подписчика
	sub, err := subscriber.Subscribe(db, cache) // делаем подписку, а также внутри обновляем БД и кэш
	if err != nil {
		fmt.Println(err)
	}
	defer sub.Close()

	s := new(models.Server) // делаем сервер и хэндлим запросы
	http.HandleFunc("/orders/",
		func(w http.ResponseWriter, r *http.Request) {
			orders(w, r, db, cache)
		})

	s.Run("8080") // запускаем сервер
}
