package conf

import (
	"os"
	"path/filepath"

	"github.com/triole/logseal"
	yaml "gopkg.in/yaml.v3"
)

type Conf struct {
	ConfigFile string
	Bind       string
	Hosts      map[string]string
	Lg         logseal.Logseal
}

func Init(configFile, bind string, lg logseal.Logseal) (conf Conf, err error) {
	var cf string
	if configFile != "" {

		cf, err = filepath.Abs(configFile)
		lg.IfErrFatal(
			"unable to construct full path",
			logseal.F{"path": configFile, "error": err},
		)

		_, err = os.Stat(configFile)
		lg.IfErrFatal(
			"file does not exist",
			logseal.F{"path": configFile, "error": err},
		)

		var raw []byte
		raw, err = os.ReadFile(configFile)
		if err != nil {
			lg.Error(
				"can not read config",
				logseal.F{"path": configFile, "error": err},
			)
		}
		err = yaml.Unmarshal(raw, &conf.Hosts)
		if err != nil {
			lg.Error(
				"can not unmarshal conf",
				logseal.F{"path": configFile, "error": err},
			)
		}
	}
	conf.Bind = bind
	conf.Lg = lg
	conf.ConfigFile = cf
	return
}
