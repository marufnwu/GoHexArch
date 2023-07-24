package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marufnwu/HexaArchProject/services"
)

func Start() {
	mux := mux.NewRouter()
	h := MyHandler{services.NewService()}

	mux.HandleFunc("/allusers", h.getAllUser).Methods(http.MethodGet)
	log.Print("Server started")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))

}