package main

import (
	"ae/src/ae"
	"ae/src/conf"
	"fmt"
	"os"

	"github.com/triole/logseal"
)

func main() {
	parseArgs()
	lg := logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	conf := conf.Init(CLI.HostnameConfig, CLI.Bind, CLI.Info, CLI.List, CLI.EnableVendors, lg)
	ae := ae.Init(conf, lg)

	if CLI.Info != "" {
		ae.PrintMacInformation(CLI.Info)
		os.Exit(0)
	}

	if CLI.List {
		ae.ListVendors()
		os.Exit(0)
	}

	if CLI.Server {
		lg.Info("run "+appName, logseal.F{"bind": CLI.Bind, "config_file": conf.ConfigFile})
		lg.Debug("configuration", logseal.F{"config": fmt.Sprintf("%+v", conf)})
		ae.RunServer()
	} else {
		ae.PrintArpTable()
	}
}
