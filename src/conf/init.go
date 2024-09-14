package conf

import (
	"os"

	"github.com/triole/logseal"
	yaml "gopkg.in/yaml.v3"
)

type Conf struct {
	Bind  string
	Hosts map[string]string
	Lg    logseal.Logseal
}

func Init(configFile, bind string, lg logseal.Logseal) (conf Conf) {
	if configFile != "" {
		var err error
		raw, err := os.ReadFile(configFile)
		if err != nil {
			lg.Error(
				"can not read config",
				logseal.F{"config": configFile, "error": err},
			)
		}
		err = yaml.Unmarshal(raw, &conf.Hosts)
		if err != nil {
			lg.Error(
				"can not unmarshal conf",
				logseal.F{"config": configFile, "error": err},
			)
		}
	}
	conf.Bind = bind
	conf.Lg = lg
	return
}
