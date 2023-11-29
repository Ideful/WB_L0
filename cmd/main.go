package main

import (
	"fmt"
	nats "l0/internal/NATS"
	cache "l0/internal/cache"
	"l0/internal/models"
	"l0/internal/repository"
	"net/http"
	"strconv"

	// service "l0/internal/service"
	"log"
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

func main() {
	var st nats.Stan
	if err := st.Connect(); err != nil {
		log.Fatal(err)
	}
	defer st.Disconnect()

	db, err := repository.CreatePostgresDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Db.Close()

	cache := cache.NewCache()
	err = cache.FillCache(db)
	if err != nil {
		fmt.Println(err)
	}

	sub, err := st.Subscribe(db, cache)
	if err != nil {
		fmt.Println(err)
	}
	defer sub.Close()

	s := new(models.Server)
	http.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		orders(w, r, db, cache)
	})

	go func() {
		s.Run("8080")
	}()
	go st.Publish()

	for {
	}
}
