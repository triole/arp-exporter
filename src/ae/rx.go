package ae

import (
	"regexp"
	"strings"
)

func (ae tAE) findAll(re *regexp.Regexp, str string) (r string) {
	rarr := re.FindAllString(str, -1)
	r = strings.Join(rarr, "")
	return
}
