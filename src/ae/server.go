package ae

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/triole/logseal"

	_ "embed"
)

//go:embed index.html
var indexPage string

func (ae tAE) RunServer() {
	http.HandleFunc("/", ae.indexPage)
	http.HandleFunc("/metrics", ae.servePrometheusMetrics)
	http.HandleFunc("/json", ae.serveJSON)
	http.ListenAndServe(ae.Conf.Bind, nil)
}

func (ae tAE) indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexPage)
}

func (ae tAE) servePrometheusMetrics(w http.ResponseWriter, r *http.Request) {
	err := ae.GetArpTable()
	if err == nil {
		metrics := ae.makePrometheusMetrics()
		ae.Lg.Debug(
			"serve prometheus metrics", logseal.F{"client": getClientIP(r)},
		)
		fmt.Fprintf(w, string(metrics))
	}
}

func (ae tAE) serveJSON(w http.ResponseWriter, r *http.Request) {
	err := ae.GetArpTable()
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		ae.Lg.Debug(
			"serve json", logseal.F{"client": getClientIP(r)},
		)
		json.NewEncoder(w).Encode(ae.ArpTable)
	}
}

func getClientIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
