package main

import (
	"log"
	"net/http"
	"sbdb-teacher/handler"
)

func main() {
	http.HandleFunc("/ping", handler.PingPongHandler)
	http.HandleFunc("/teacher", handler.Handler)
	http.HandleFunc("/teachers", handler.AllHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
