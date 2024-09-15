package ae

import (
	"encoding/json"
	"fmt"
)

func (ae *tAE) pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(s))
}
