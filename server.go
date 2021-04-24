package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

var router = mux.NewRouter()

func main()  {
	server := http.Server{Addr: ":8080"}
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello"))
	})
	http.Handle("/", router)

	_ = server.ListenAndServe()
}
