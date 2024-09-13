package main

import (
	"regexp"
	"strings"
)

type tRx struct {
	MAC *regexp.Regexp
	IP  *regexp.Regexp
}

func initRx() (rx tRx) {
	return tRx{
		MAC: regexp.MustCompile(
			"([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})",
		),
		IP: regexp.MustCompile(
			"[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}",
		),
	}
}

func (rx tRx) findAll(re *regexp.Regexp, str string) (r string) {
	rarr := re.FindAllString(str, -1)
	r = strings.Join(rarr, "")
	return
}
