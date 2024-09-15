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
	conf := conf.Init(CLI.Config, CLI.Bind, CLI.EnableVendors, lg)
	ae := ae.Init(conf, lg)

	if CLI.Server {
		lg.Info("run "+appName, logseal.F{"bind": CLI.Bind, "config_file": conf.ConfigFile})
		lg.Debug("configuration", logseal.F{"config": fmt.Sprintf("%+v", conf)})
		ae.RunServer()
	} else {
		ae.PrintArpTable()
	}
}
