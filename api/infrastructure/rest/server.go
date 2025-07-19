package rest

import (
	"log"
	"loom-api/api/infrastructure/rest/poc"
	"net/http"
)

func StartServer(addr string, handler *poc.Handler) error {
	http.HandleFunc("/pocs", handler.CreatePocHandler)
	log.Printf("API REST corriendo en %s", addr)
	return http.ListenAndServe(addr, nil)
}
