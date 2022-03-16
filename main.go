package main

import (
	"net/http"

	"encurtUrl/db"
	"encurtUrl/url"
)

func main() {

	db.Connection()
	defer db.DB.Close()
	http.HandleFunc("/", url.CreateEndPoint)
	http.ListenAndServe(":8080", nil)
}
