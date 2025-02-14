package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	reverseproxy "github.com/rafa-souza-dev/dog-balancer/reverse_proxy"
)

var (
	backends = []string{
		"http://localhost:8001",
		"http://localhost:8002",
		"http://localhost:8003",
		"http://localhost:8004",
	}

	currentIndex = 0
)

func main() {
	log.Println("Dog-balancer started!")
	mux := http.NewServeMux()
	mu := sync.Mutex{}

	mux.HandleFunc("/", reverseproxy.HandleRedirectRequest(
		&currentIndex,
		&mu,
		backends,
	))

	server := &http.Server{
		Addr: ":8000",
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
