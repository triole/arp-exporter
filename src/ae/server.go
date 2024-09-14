package ae

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/triole/logseal"
)

func (ae tAE) RunServer() {
	http.HandleFunc("/metrics", ae.servePrometheusMetrics)
	http.HandleFunc("/json", ae.ServeJSON)
	http.ListenAndServe(ae.Conf.Bind, nil)
}

func (ae tAE) servePrometheusMetrics(w http.ResponseWriter, r *http.Request) {
	err := ae.GetArpTable()
	if err == nil {
		metrics := ""
		for _, el := range ae.ArpTable {
			metrics += "#HELP\n"
			metrics += "#TYPE\n"
			metrics += "arp_exporter{" +
				"ip=\"" + el.IP + "\", mac=\"" + el.MAC +
				"\"} 0\n"
		}
		ae.Lg.Debug(
			"serve prometheus metrics", logseal.F{"client": getClientIP(r)},
		)
		fmt.Fprintf(w, string(metrics))
	}
}

func (ae tAE) ServeJSON(w http.ResponseWriter, r *http.Request) {
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
