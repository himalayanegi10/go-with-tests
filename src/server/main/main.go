package main

import (
	"log"
	"net/http"
	"github.com/himalayanegi10/go-with-tests/src/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}