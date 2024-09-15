package ae

import (
	"os/exec"
	"strings"

	"github.com/triole/logseal"
)

func (ae *tAE) GetArpTable() (err error) {
	var by []byte
	ae.ArpTable = []tArpEntry{}
	by, err = exec.Command("arp", "-an").Output()
	if err == nil {
		for _, line := range strings.Split(string(by), "\n") {
			newEntry := tArpEntry{
				MAC: ae.findAll(ae.Rx.MAC, line),
				IP:  ae.findAll(ae.Rx.IP, line),
			}
			if newEntry.MAC != "" && newEntry.IP != "" {
				newEntry.Name = ae.Conf.GetHostName(newEntry.MAC)
				var vendor tVendor
				if ae.Conf.EnableVendors {
					vendor = ae.getVendor(newEntry.MAC)
				}
				if vendor.Name != "" {
					newEntry.Vendor = vendor.Name
				}
				ae.ArpTable = append(ae.ArpTable, newEntry)
			}
		}
	} else {
		ae.Lg.Error("unable to get arp table", logseal.F{"error": err})
	}
	return
}

func (ae *tAE) PrintArpTable() {
	err := ae.GetArpTable()
	if err == nil {
		ae.pprint(ae.ArpTable)
	}
}
