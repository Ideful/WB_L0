package models

import (

	// "l0/internal/repository"
	"net/http"
	// "strconv"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// func id(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.Write(loginFormTmpl)
// 		return
// 	}

// 	// r.ParseForm()
// 	// inputLogin := r.Form["login"][0]

// 	inputLogin := r.FormValue("login")
// 	fmt.Fprintln(w, "you enter: ", inputLogin)
// }

type IDResponse struct {
	ID string `json:"id"`
}

func (s *Server) Run(port string) error {
	// someValue, ok := ctx.Value("someKey").(string)
	// if ok {
	// 	fmt.Println("Value from context:", someValue)
	// } else {
	// 	fmt.Println("Value not found in context")
	// }

	s.httpServer = &http.Server{
		Addr: ":" + port,
		// Handler:        handler,
		ReadTimeout:    4 * time.Second,
		WriteTimeout:   4 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// v := <-s.StringChannel
	return s.httpServer.ListenAndServe()
}
