package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Dog-balancer started!")
	mux := http.NewServeMux()

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 1 * time.Minute,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
		panic(err)
	}
}
