package ae

import (
	"net"
	"strings"

	"github.com/triole/logseal"
)

func (ae *tAE) appendIP(params tParams, addr string) (r bool) {
	var b bool
	var err error
	if len(params.Include) == 0 {
		r = true
	} else {
		b, err := ae.cidrRangeContains(params.Include, addr)
		if b && err == nil {
			r = true
		}
		if err != nil {
			r = true
		}
	}
	b, err = ae.cidrRangeContains(params.Exclude, addr)
	if b && err == nil {
		r = false
	}
	ae.Lg.Debug(
		"append ip",
		logseal.F{"ip": addr, "excl": params.Exclude, "incl": params.Include, "result": b, "error": err},
	)
	return
}

func (ae *tAE) cidrRangeContains(cidrRanges []string, addr string) (bool, error) {
	ip := net.ParseIP(addr)
	for _, cidr := range cidrRanges {
		i := strings.IndexByte(cidr, '/')
		if i < 0 {
			pIP := net.ParseIP(cidr)
			if pIP == nil {
				return false, nil
			}
			if pIP.Equal(ip) {
				return true, nil
			}
		} else {
			_, nets, err := net.ParseCIDR(cidr)
			if err != nil {
				ae.Lg.Warn("can not parse cidr range, skip to apply filter", logseal.F{"cidr": cidr})
				return false, err
			}
			if nets.Contains(ip) {
				return true, nil
			}
		}
		if ip == nil {
			return false, nil
		}
	}
	return false, nil
}
