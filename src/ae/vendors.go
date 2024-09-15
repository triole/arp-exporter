package ae

import (
	_ "embed"
	"encoding/json"
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
