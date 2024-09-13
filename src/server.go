package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func runServer() {
	http.HandleFunc("/metrics", serveMetrics)
	http.HandleFunc("/json", serveJSON)
	http.ListenAndServe(CLI.Bind, nil)
}

func serveMetrics(w http.ResponseWriter, r *http.Request) {
	arp, _ := initArpTable()
	fmt.Fprintf(w, string(arp.metrics()))
}

func serveJSON(w http.ResponseWriter, r *http.Request) {
	arp, _ := initArpTable()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(arp)
}
