package main

import (
	"log"
	"net/http"
	"wxcloudrun-golang/service"
)

func main() {
	http.HandleFunc("/api/rotime", service.TimeHandler)

	log.Fatal(http.ListenAndServe(":8851", nil))
}
