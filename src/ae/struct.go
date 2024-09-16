package ae

import (
	"ae/src/conf"
	"bytes"
	"net"
	"regexp"

	"github.com/triole/logseal"
)

type tAE struct {
	ArpTable  tArpTable
	Hostnames map[string]string
	Conf      conf.Conf
	Rx        tRegexSchemes
	Vendors   tVendorsMap
	Lg        logseal.Logseal
}

type tArpTable []tArpEntry
type tArpEntry struct {
	MAC    string `json:"mac"`
	IP     string `json:"ip"`
	Itf    string `json:"itf,omitempty"`
	Name   string `json:"name,omitempty"`
	Vendor string `json:"vendor,omitempty"`
}

type tRegexSchemes struct {
	MAC *regexp.Regexp
	IP  *regexp.Regexp
}

func (arr tArpTable) Len() int {
	return len(arr)
}

func (arr tArpTable) Less(i, j int) bool {
	s1 := net.ParseIP(arr[i].IP)
	s2 := net.ParseIP(arr[j].IP)
	return bytes.Compare(s1, s2) < 0
}

func (arr tArpTable) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
