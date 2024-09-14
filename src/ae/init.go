package ae

import (
	"ae/src/conf"
	"regexp"

	"github.com/triole/logseal"
)

type tAE struct {
	ArpTable  []tArpEntry
	Hostnames map[string]string
	Conf      conf.Conf
	Rx        tRegexSchemes
	Lg        logseal.Logseal
}

type tArpEntry struct {
	MAC string `json:"mac"`
	IP  string `json:"ip"`
}

type tRegexSchemes struct {
	MAC *regexp.Regexp
	IP  *regexp.Regexp
}

func Init(conf conf.Conf, lg logseal.Logseal) (ae tAE) {
	ae.Conf = conf
	ae.Rx = initRegexSchemes()
	ae.Lg = lg
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
