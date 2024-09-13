package main

import (
	"regexp"
	"strings"
)

func findAll(rx string, str string) (r string) {
	re := regexp.MustCompile(rx)
	rarr := re.FindAllString(str, -1)
	r = strings.Join(rarr, "")
	return
}
