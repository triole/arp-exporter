package main

import (
	"io"
	"net/http"
)

type tResponse struct {
	Method  string
	Proto   string
	Host    string
	URL     string
	Request tRequest
}

type tRequest struct {
	Params  map[string][]string
	Body    string
	Headers map[string][]string
}

type handler struct{}

func (h *handler) serveHTTP(resp http.ResponseWriter, req *http.Request) {
	// arp := getArpTable()
	// lg.LogInfo("Got request", logrus.Fields{"data": string(jsonData)})
	resp.Header().Set("Content-Type", "application/json")
	io.WriteString(resp, "This is my website!\n")
}
