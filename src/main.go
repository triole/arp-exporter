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
	conf, err := conf.Init(CLI.Config, CLI.Bind, lg)
	ae := ae.Init(conf, lg)

	if err == nil {
		lg.Info("run "+appName, logseal.F{"bind": CLI.Bind, "config_file": conf.ConfigFile})
	} else {
		lg.Info("run "+appName, logseal.F{"bind": CLI.Bind})
	}
	lg.Debug("configuration", logseal.F{"config": fmt.Sprintf("%+v", conf)})
	if CLI.Print {
		err := ae.GetArpTable()
		if err == nil {
			fmt.Printf("%+v\n", ae.ArpTable)
		}
	} else {
		ae.RunServer()
	}
}
