package ae

import (
	"ae/src/conf"
	"regexp"

	"github.com/triole/logseal"

	_ "embed"
)

func Init(conf conf.Conf, lg logseal.Logseal) (ae tAE) {
	ae.Conf = conf
	ae.Rx = initRegexSchemes()
	ae.Lg = lg
	if ae.Conf.Info != "" || ae.Conf.List || ae.Conf.EnableVendors {
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
