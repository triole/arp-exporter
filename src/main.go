package main

import (
	"fmt"
	"net/http"

	"github.com/triole/logseal"
)

func main() {
	parseArgs()
	lg := logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)

	lg.Info("run "+appName, logseal.F{"bind": CLI.Bind})
	http.HandleFunc("/", serveMetrics)
	http.ListenAndServe(CLI.Bind, nil)
}

func serveMetrics(w http.ResponseWriter, r *http.Request) {
	arp, _ := getArpTable()
	fmt.Fprintf(w, string(arp))
}
