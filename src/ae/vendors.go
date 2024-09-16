package ae

import (
	_ "embed"
	"encoding/json"
	"errors"
	"strings"

	"github.com/triole/logseal"
)

//go:embed vendors.json
var vendorsJSON string

type tVendorsList []tVendor
type tVendorsMap map[string]tVendor

type tVendor struct {
	MacPrefix  string `json:"macPrefix"`
	Name       string `json:"vendorName"`
	Private    bool   `json:"private"`
	BlockType  string `json:"blockType"`
	LastUpdate string `json:"lastUpdate"`
}

func (ae *tAE) PrintMacInformation(mac string) {
	var vendor tVendor
	if len(mac) >= 8 {
		vendor = ae.getVendor(mac)
	} else {
		err := errors.New("invalid mac or mac prefix: " + mac)
		ae.Lg.Error("can not look up mac address", logseal.F{"error": err})
	}
	if vendor.Name == "" {
		ae.Lg.Info("no vendor information found", logseal.F{"mac": mac})
	} else {
		ae.pprint(vendor)
	}
}

func (ae *tAE) unmarshalVendors() {
	var vendorList tVendorsList
	err := json.Unmarshal([]byte(vendorsJSON), &vendorList)
	if err != nil {
		ae.Lg.Fatal(
			"error during unmarshal vendor json: ",
			logseal.F{"error": err},
		)
	}
	ae.Vendors = make(map[string]tVendor)
	for _, ven := range vendorList {
		ae.Vendors[strings.ToLower(ven.MacPrefix)] = ven
	}
}

func (ae tAE) getVendor(mac string) (vendor tVendor) {
	var prefix string
	for i := 12; i >= 1; i-- {
		prefix = strings.ToLower(mac[0:i])
		if val, ok := ae.Vendors[prefix]; ok {
			vendor = val
			break
		}
	}
	return
}
