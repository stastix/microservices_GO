package main

import (
	"log"
	"net/http"
	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("HTTP server erro : %v", err)
	}

}
