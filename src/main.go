package main

import (
	"ae/src/ae"
	"ae/src/conf"
	"fmt"

	"github.com/triole/logseal"
)

func main() {
	parseArgs()
	lg := logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	conf := conf.Init(CLI.Config, CLI.Bind, lg)
	ae := ae.Init(conf, lg)

	lg.Info("run "+appName, logseal.F{"bind": CLI.Bind})
	if CLI.Print {
		err := ae.GetArpTable()
		if err == nil {
			fmt.Printf("%+v\n", ae.ArpTable)
		}
	} else {
		ae.RunServer()
	}
}
