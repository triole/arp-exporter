package ae

import (
	"ae/src/conf"
	"testing"

	"github.com/triole/logseal"
)

func TestAppendIP(t *testing.T) {
	validateAppendIP(
		newParams([]string{"127.0.0.0/28"}, []string{"127.0.0.4"}),
		"127.0.0.2", true, t,
	)
	validateAppendIP(
		newParams([]string{"127.0.0.0/28"}, []string{"127.0.0.4"}),
		"127.0.0.4", false, t,
	)
	validateAppendIP(
		newParams([]string{"127.0.0.0/28"}, []string{"127.0.0.4"}),
		"127.0.0.44", false, t,
	)
}

func validateAppendIP(params tParams, ip string, exp bool, t *testing.T) {
	var conf conf.Conf
	lg := logseal.Init()
	ae := Init(conf, lg)
	res := ae.appendIP(params, ip)
	if exp != res {
		t.Errorf(
			"appendIP failed exp!=res: %v!=%v, params: %+v, ip: %s",
			exp, res, params, ip,
		)
	}
}

func TestIsIPinRange(t *testing.T) {
	validateCidrRangeContains([]string{}, "127.0.0.4", false, t)
	validateCidrRangeContains([]string{"127.0.0.0/28"}, "127.0.0.4", true, t)
	validateCidrRangeContains([]string{"127.0.0.0/28"}, "127.0.0.24", false, t)
	validateCidrRangeContains([]string{}, "INVALID", false, t)
}

func validateCidrRangeContains(cidr []string, ip string, exp bool, t *testing.T) {
	var conf conf.Conf
	var lg logseal.Logseal
	ae := Init(conf, lg)
	res, _ := ae.cidrRangeContains(cidr, ip)
	if exp != res {
		t.Errorf(
			"cidrRangeContains failed exp!=res: %v!=%v, params: %+v, ip: %s",
			exp, res, cidr, ip,
		)
	}
}

func newParams(incl, excl []string) (params tParams) {
	params.Include = incl
	params.Exclude = excl
	return
}
