package ae

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/triole/logseal"
)

//go:embed vendors.json
var vendorsJSON string

type tVendors []tVendor

type tVendor struct {
	MacPrefix  string `json:"macPrefix"`
	Name       string `json:"vendorName"`
	Private    bool   `json:"private"`
	BlockType  string `json:"blockType"`
	LastUpdate string `json:"lastUpdate"`
}

func (arr tVendors) Len() int {
	return len(arr)
}

func (arr tVendors) Less(i, j int) bool {
	return arr[i].MacPrefix < arr[j].MacPrefix
}

func (arr tVendors) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (ae *tAE) ListVendors() {
	sort.Sort(ae.Vendors)
	for _, el := range ae.Vendors {
		fmt.Printf("%-16s %s\n", el.MacPrefix, el.Name)
	}
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
	err := json.Unmarshal([]byte(vendorsJSON), &ae.Vendors)
	if err != nil {
		ae.Lg.Fatal(
			"error during unmarshal vendor json: ",
			logseal.F{"error": err},
		)
	}
}

func (ae tAE) getVendor(mac string) (vendor tVendor) {
	prefix := mac[0:8]
	for _, el := range ae.Vendors {
		if strings.EqualFold(el.MacPrefix, prefix) {
			vendor = el
			break
		}
	}
	return
}
