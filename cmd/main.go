package main

import (
	"fmt"
	nats "l0/internal/NATS"
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
	<form action="/id/" method="get">
		<input type="text" name="id">
		<input type="submit" value="id">
	</form>
	</body>
</html>
`)

func id(w http.ResponseWriter, r *http.Request, db *repository.MyDB) {
	w.Write(loginFormTmpl)
	if r.Method != http.MethodGet {
		return
	}

	id := r.FormValue("id")
	val, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	o, err := db.GetOrder(val)
	if err != nil {
		return
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

	sub, err := st.Subscribe(db)
	if err != nil {
		fmt.Println(err)
	}
	defer sub.Close()

	s := new(models.Server)
	http.HandleFunc("/id/", func(w http.ResponseWriter, r *http.Request) {
		id(w, r, db)
	})

	go func() {
		s.Run("8080")
	}()
	// go st.Publish()

	for {
	}
}
