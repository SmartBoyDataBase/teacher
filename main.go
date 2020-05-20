package main

import (
	"log"
	"net/http"
	"os"
	"sbdb-teacher/handler"
)

func main() {
	http.HandleFunc("/ping", handler.PingPongHandler)
	http.HandleFunc("/teacher", handler.Handler)
	http.HandleFunc("/teachers", handler.AllHandler)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
