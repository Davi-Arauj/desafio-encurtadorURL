package main

import (
	
	"net/http"
	"strings"

	"encurtUrl/db"
	"encurtUrl/url"
)

func main() {

	db.Connection()
	
	configurandoServidor()

	defer db.DB.Close()

}

func configurandoServidor() {
	
	configurarRotas()
	http.ListenAndServe(":8080", nil)

}

func configurarRotas() {
	
	http.HandleFunc("/create", direcionamentoRotas)
	http.HandleFunc("/url/", direcionamentoRotas)
	http.HandleFunc("/redirect/", direcionamentoRotas)
}

func direcionamentoRotas(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")

	if r.Method == "POST" {
		url.CreateEndPoint(w, r)
	} else if r.Method == "GET" && partes[1] == "url" {
		url.ExpandEndPoint(w, r)
	} else if r.Method == "GET" && partes[1] == "redirect" {
		url.RootEndPoint(w, r)
	}

}
