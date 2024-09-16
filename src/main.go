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
	conf := conf.Init(CLI, lg)
	ae := ae.Init(conf, lg)

	if CLI.MacInfo != "" {
		ae.PrintMacInformation(CLI.MacInfo)
		os.Exit(0)
	}

	if CLI.ListVendors {
		// ae.ListVendors()
		os.Exit(0)
	}

	if CLI.Server {
		fields := logseal.F{}
		fields["bind"] = CLI.Bind
		if conf.ArpTable != "" {
			fields["arptable_file"] = conf.ArpTable
		}
		if conf.HostnameConfig != "" {
			fields["hostname_config"] = conf.HostnameConfig
		}
		lg.Info("run "+appName, fields)
		lg.Debug("configuration", logseal.F{"config": fmt.Sprintf("%+v", conf)})
		ae.RunServer()
	} else {
		ae.PrintArpTable()
	}
}
