package main

import (
	"os/exec"
	"strings"
)

func getArpTable() (arpList string, err error) {
	var by []byte
	by, err = exec.Command("arp", "-an").Output()
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(by), "\n") {
		mac := findAll("([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})", line)
		ip := findAll("[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}", line)
		if mac != "" && ip != "" {
			arpList += "#HELP\n"
			arpList += "#TYPE\n"
			arpList += "arp_exporter{" +
				"ip=\"" + ip + "\", mac=\"" + mac +
				"\"} 0\n"
		}
	}
	return
}
