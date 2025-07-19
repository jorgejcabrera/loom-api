package rest

import (
	"log"
	"net/http"
)

func StartServer(addr string, router http.Handler) error {
	log.Printf("Running loom on %s", addr)
	return http.ListenAndServe(addr, router)
}
