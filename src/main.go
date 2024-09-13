package main

import (
	"fmt"

	"github.com/triole/logseal"
)

var (
	lg logseal.Logseal
	rx tRx
)

func main() {
	parseArgs()
	rx = initRx()
	lg = logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)

	lg.Info("run "+appName, logseal.F{"bind": CLI.Bind})
	if CLI.Print {
		arp, _ := initArpTable()
		fmt.Printf("%+v\n", arp)
	} else {
		runServer()
	}
}
