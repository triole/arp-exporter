package main

import (
	"os/exec"
	"strings"
)

type tArpTable []tArp

type tArp struct {
	MAC string `json:"mac"`
	IP  string `json:"ip"`
}

func initArpTable() (arpList tArpTable, err error) {
	var by []byte
	by, err = exec.Command("arp", "-an").Output()
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(by), "\n") {
		arp := tArp{
			MAC: findAll("([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})", line),
			IP:  findAll("[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}", line),
		}
		if arp.MAC != "" && arp.IP != "" {
			arpList = append(arpList, arp)
		}
	}
	return
}

func (arp tArpTable) metrics() (metrics string) {
	for _, el := range arp {
		metrics += "#HELP\n"
		metrics += "#TYPE\n"
		metrics += "arp_exporter{" +
			"ip=\"" + el.IP + "\", mac=\"" + el.MAC +
			"\"} 0\n"
	}
	return
}
