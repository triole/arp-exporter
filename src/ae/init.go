package ae

import (
	"ae/src/conf"
	"regexp"

	"github.com/triole/logseal"

	_ "embed"
)

type tAE struct {
	ArpTable  []tArpEntry
	Hostnames map[string]string
	Conf      conf.Conf
	Rx        tRegexSchemes
	Vendors   tVendors
	Lg        logseal.Logseal
}

type tArpEntry struct {
	MAC    string `json:"mac"`
	IP     string `json:"ip"`
	Name   string `json:"name,omitempty"`
	Vendor string `json:"vendor,omitempty"`
}

type tRegexSchemes struct {
	MAC *regexp.Regexp
	IP  *regexp.Regexp
}

func Init(conf conf.Conf, lg logseal.Logseal) (ae tAE) {
	ae.Conf = conf
	ae.Rx = initRegexSchemes()
	ae.Lg = lg
	if ae.Conf.EnableVendors {
		ae.unmarshalVendors()
	}
	return
}

func initRegexSchemes() (rx tRegexSchemes) {
	return tRegexSchemes{
		MAC: regexp.MustCompile(
			"([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})",
		),
		IP: regexp.MustCompile(
			"[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}",
		),
	}
}
