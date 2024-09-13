package main

import (
	"fmt"
	"net/http"
)

func runServer() {
	http.HandleFunc("/", serveMetrics)
	http.ListenAndServe(CLI.Bind, nil)
}

func serveMetrics(w http.ResponseWriter, r *http.Request) {
	arp, _ := getArpTable()
	fmt.Fprintf(w, string(fmt.Sprintf("%s", arp)))
}
