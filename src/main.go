package main

import (
	"fmt"

	"github.com/triole/logseal"
)

func main() {
	parseArgs()
	lg := logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)

	lg.Info("run "+appName, logseal.F{"bind": CLI.Bind})
	if CLI.Print {
		arp, _ := getArpTable()
		fmt.Printf("%+v\n", arp)
	} else {
		runServer()
	}
}
